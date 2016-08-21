package tmx

import (
	"encoding/xml"
	"io/ioutil"
)

// Load loads a TMX map from a file stored at the specified path.
// If the loading fails, it returns an error.
func Load(path string) (*Map, error) {
	tmxContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tmxMap := new(Map)
	err = xml.Unmarshal(tmxContent, tmxMap)
	if err != nil {
		return nil, err
	}

	for i := range tmxMap.Tilesets {
		if tmxMap.Tilesets[i].Source != "" {
			tilesetContent, err := ioutil.ReadFile(tmxMap.Tilesets[i].Source)
			if err != nil {
				return nil, err
			}

			var tileset Tileset
			err = xml.Unmarshal(tilesetContent, &tileset)
			if err != nil {
				return nil, err
			}

			tmxMap.Tilesets[i] = tileset
		}
	}

	return tmxMap, nil
}

// Map is a level map in TMX format.
type Map struct {
	XMLName xml.Name `xml:"map"`

	Version         string `xml:"version,attr"`
	RenderOrder     string `xml:"renderorder,attr"`
	Width           int    `xml:"width,attr"`
	Height          int    `xml:"height,attr"`
	TileWidth       int    `xml:"tilewidth,attr"`
	TileHeight      int    `xml:"tileheight,attr"`
	BackgroundColor string `xml:"backgroundcolor,attr"`
	NextObjectID    int    `xml:"nextobjectid,attr"`

	Props Properties `xml:"properties"`

	Tilesets     []Tileset     `xml:"tileset"`
	Layers       []Layer       `xml:"layer"`
	ObjectGroups []ObjectGroup `xml:"objectgroup"`
	ImageLayers  []ImageLayer  `xml:"imagelayer"`
}

// Properties is a list of properties.
type Properties struct {
	List []Property `xml:"property"`
}

// Property is an arbitrary property. It has an arbitrary name, type and value.
type Property struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

// Tileset is an image grid of tiles.
type Tileset struct {
	XMLName xml.Name `xml:"tileset"`

	FirstGID   int    `xml:"firstgid,attr"`
	Source     string `xml:"source,attr"`
	Name       string `xml:"name,attr"`
	TileWidth  int    `xml:"tilewidth,attr"`
	TileHeight int    `xml:"tileheight,attr"`
	Spacing    int    `xml:"spacing,attr"`
	Margin     int    `xml:"margin,attr"`
	TileCount  int    `xml:"tilecount,attr"`
	Columns    int    `xml:"columns,attr"`

	Props Properties `xml:"properties"`

	TileOffset struct {
		X int `xml:"x,attr"`
		Y int `xml:"y,attr"`
	} `xml:"tileoffset"`

	Image Image `xml:"image"`
}

// Layer is a single layout of tiles.
type Layer struct {
	Name    string  `xml:"name,attr"`
	Opacity float64 `xml:"opacity,attr"`
	Visible bool    `xml:"visible,attr"`
	OffsetX int     `xml:"offsetx,attr"`
	OffsetY int     `xml:"offsety,attr"`

	Props Properties `xml:"properties"`

	Data struct {
		Encoding    string `xml:"encoding,attr"`
		Compression string `xml:"compression,attr"`
		Content     string `xml:",chardata"`
	} `xml:"data"`
}

// UnmarshalXML for default values for fuck sake.
func (l *Layer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type layer Layer
	defaultLayer := layer{
		Opacity: 1.0,
		Visible: true,
	}
	if err := d.DecodeElement(&defaultLayer, &start); err != nil {
		return err
	}
	*l = Layer(defaultLayer)
	return nil
}

// ObjectGroup is a single list of objects.
type ObjectGroup struct {
	Name    string  `xml:"name,attr"`
	Color   string  `xml:"color,attr"`
	Opacity float64 `xml:"opacity,attr"`
	Visible bool    `xml:"visible,attr"`
	OffsetX int     `xml:"offsetx,attr"`
	OffsetY int     `xml:"offsety,attr"`

	Props Properties `xml:"properties"`

	Objects []Object `xml:"object"`
}

// UnmarshalXML for default values for fuck sake.
func (o *ObjectGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type objectGroup ObjectGroup
	defaultGroup := objectGroup{
		Opacity: 1.0,
		Visible: true,
	}
	if err := d.DecodeElement(&defaultGroup, &start); err != nil {
		return err
	}
	*o = ObjectGroup(defaultGroup)
	return nil
}

// Object is an arbitrary game object.
type Object struct {
	ID       int    `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr"`
	X        int    `xml:"x,attr"`
	Y        int    `xml:"y,attr"`
	Width    int    `xml:"width,attr"`
	Height   int    `xml:"height,attr"`
	Rotation int    `xml:"rotation,attr"`
	Visible  bool   `xml:"visible,attr"`

	Props Properties `xml:"properties"`
}

// ImageLayer is a single layout of images.
type ImageLayer struct {
	Name    string  `xml:"name,attr"`
	OffsetX int     `xml:"offsetx,attr"`
	OffsetY int     `xml:"offsety,attr"`
	Opacity float64 `xml:"opacity,attr"`
	Visible bool    `xml:"visible,attr"`

	Props Properties `xml:"properties"`

	Images []Image `xml:"image"`
}

// UnmarshalXML for default values for fuck sake.
func (i *ImageLayer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type imageLayer ImageLayer
	defaultLayer := imageLayer{
		Opacity: 1.0,
		Visible: true,
	}
	if err := d.DecodeElement(&defaultLayer, &start); err != nil {
		return err
	}
	*i = ImageLayer(defaultLayer)
	return nil
}

// Image is a picture file.
type Image struct {
	Format string `xml:"format,attr"`
	Source string `xml:"source,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}
