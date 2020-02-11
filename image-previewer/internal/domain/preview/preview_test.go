package preview

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"gitlab.com/tirava/image-previewer/internal/caches"
	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/encoders"
	"gitlab.com/tirava/image-previewer/internal/models"
	"gitlab.com/tirava/image-previewer/internal/previewers"
	"gitlab.com/tirava/image-previewer/internal/storages"
)

// nolint:gochecknoglobals
var testCases = []struct {
	description   string
	originalImage string
	previews      [][2]int
	previewer     string
}{
	{
		"JPG: width > height",
		"_gopher_original_1024x504.jpg",
		[][2]int{
			{1024, 252},
			{2000, 1000},
			{200, 700},
			{256, 126},
			{333, 666},
			{500, 500},
			{50, 50},
		},
		"nfnt_crop",
	},
	{
		"PNG: width < height",
		"_gopher_original_540x720.png",
		[][2]int{
			{540, 360},
			{1000, 1500},
			{130, 1000},
			{130, 180},
			{175, 850},
			{500, 500},
			{50, 50},
		},
		"xdraw",
	},
}

// nolint:wsl
func TestPreview(t *testing.T) {
	for _, test := range testCases {
		conf := models.Config{
			Previewer:       test.previewer,
			ImageURLEncoder: "md5",
			Cacher:          "nolimit",
			MaxCacheItems:   0,
			Storager:        "inmemory",
		}
		prev, err := initPreview(conf)
		if err != nil {
			t.Fatal(err)
		}

		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			ext := filepath.Ext(test.originalImage)

			fileName := filepath.Join("../", "../", "../", "examples", test.originalImage)
			file, err := os.Open(fileName)

			if err != nil {
				t.Fatal(err)
			}

			var img image.Image

			switch ext {
			case ".jpg", ".jpeg":
				img, err = jpeg.Decode(file)
			case ".png":
				img, err = png.Decode(file)
			case ".gif":
				img, err = gif.Decode(file)
			}

			if err != nil {
				t.Fatal(err)
			}

			for _, pr := range test.previews {
				r := prev.Preview(pr[0], pr[1], img, entities.ResizeOptions{})

				w := r.Bounds().Max.X - r.Bounds().Min.X
				h := r.Bounds().Max.Y - r.Bounds().Min.Y

				if w != pr[0] || h != pr[1] {
					t.Errorf("'%s' preview bounds expected - %dx%d\nbut resized to - %dx%d\n",
						test.originalImage, pr[0], pr[1], w, h)
				}
			}

			file.Close()
		})
	}
}

const (
	benchMaxX   = 250
	benchMaxY   = 250
	benchWidth  = 200
	benchHeight = 200
)

func benchRGBA(b *testing.B, prev *Preview, opts entities.ResizeOptions) {
	m := image.NewRGBA(image.Rect(0, 0, benchMaxX, benchMaxY))

	for y := m.Rect.Min.Y; y < m.Rect.Max.Y; y++ {
		for x := m.Rect.Min.X; x < m.Rect.Max.X; x++ {
			i := m.PixOffset(x, y)
			m.Pix[i+0] = uint8(y + 4*x)
			m.Pix[i+1] = uint8(y + 4*x)
			m.Pix[i+2] = uint8(y + 4*x)
			m.Pix[i+3] = uint8(4*y + x)
		}
	}

	var out image.Image

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out = prev.Preview(benchWidth, benchHeight, m, opts)
	}

	out.At(0, 0)
}

func Benchmark_XDraw_Nearest_RGBA(b *testing.B) {
	conf := models.Config{
		Previewer:       "xdraw",
		ImageURLEncoder: "md5",
		Cacher:          "nolimit",
		MaxCacheItems:   0,
		Storager:        "inmemory",
	}
	prev, _ := initPreview(conf)

	benchRGBA(b, prev, entities.ResizeOptions{})
}

func Benchmark_Nfnt_Nearest_RGBA(b *testing.B) {
	conf := models.Config{
		Previewer:       "nfnt_crop",
		ImageURLEncoder: "md5",
		Cacher:          "nolimit",
		MaxCacheItems:   0,
		Storager:        "inmemory",
	}
	prev, _ := initPreview(conf)

	benchRGBA(b, prev, entities.ResizeOptions{})
}

func initPreview(conf models.Config) (*Preview, error) {
	prev, err := previewers.NewPreviewer(conf.Previewer)
	if err != nil {
		return nil, err
	}

	enc, err := encoders.NewImageURLEncoder(conf.ImageURLEncoder)
	if err != nil {
		return nil, err
	}

	stor, err := storages.NewStorager(conf.Storager, conf.StoragePath)
	if err != nil {
		return nil, err
	}

	cash, err := caches.NewCacher(conf.Cacher, stor, conf.MaxCacheItems)
	if err != nil {
		return nil, err
	}

	return NewPreview(prev, enc, cash, stor)
}
