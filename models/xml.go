// Package models contains all the necessary type declarations required for the Domain objects of the assignment
package models

import "encoding/xml"

// XML is the root level of XML to be read.
type XML struct {
	Works xml.Name `xml:"works>work"`
	Work  []Work   `xml:"work"`
}

// Work is definition of single piece of work/image
type Work struct {
	ID          int    `xml:"id"`
	Filename    string `xml:"filename"`
	ImageWidth  int    `xml:"image_width"`
	ImageHeight int    `xml:"image_height"`
	URLs        []URL  `xml:"urls>url"`
	EXIF        EXIF   `xml:"exif"`
}

// EXIF contains all the metadata of the image taken
type EXIF struct {
	Flash                       string `xml:"flash"`
	Sharpness                   string `xml:"sharpness"`
	PixelYDimension             string `xml:"pixel_y_dimension"`
	DateTimeDigitized           string `xml:"date_time_digitized"`
	ExposureBiasValue           string `xml:"exposure_bias_value"`
	FocalLengthIn35mmFilm       string `xml:"focal_length_in_35mm_film"`
	SubsecTimeOrginal           string `xml:"subsec_time_orginal"`
	Model                       string `xml:"model"`
	Software                    string `xml:"software"`
	FocalLength                 string `xml:"focal_length"`
	MaxApertureValue            string `xml:"max_aperture_value"`
	SceneCaptureType            string `xml:"scene_capture_type"`
	SubsecTimeDigitized         string `xml:"subsec_time_digitized"`
	IsoSpeedRatings             string `xml:"iso_speed_ratings"`
	DateTime                    string `xml:"date_time"`
	CompressedBitsPerPixel      string `xml:"compressed_bits_per_pixel"`
	SubjectDistanceRange        string `xml:"subject_distance_range"`
	CustomRendered              string `xml:"custom_rendered"`
	Make                        string `xml:"make"`
	GainControl                 string `xml:"gain_control"`
	Orientation                 string `xml:"orientation"`
	ExposureTime                string `xml:"exposure_time"`
	ExposureMode                string `xml:"exposure_mode"`
	XResolution                 string `xml:"x_resolution"`
	MeteringMode                string `xml:"metering_mode"`
	SensingMethod               string `xml:"sensing_method"`
	Contrast                    string `xml:"contrast"`
	ColorSpace                  string `xml:"color_space"`
	FNumber                     string `xml:"f_number"`
	WhiteBalance                string `xml:"white_balance"`
	UserComment                 string `xml:"user_comment"`
	YResolution                 string `xml:"y_resolution"`
	ResolutionUnit              string `xml:"resolution_unit"`
	LightSource                 string `xml:"light_source"`
	Saturation                  string `xml:"saturation"`
	PixelXDimension             string `xml:"pixel_x_dimension"`
	DateTimeOriginal            string `xml:"date_time_original"`
	YcbCrPositioning            string `xml:"ycb_cr_positioning"`
	ExposureProgram             string `xml:"exposure_program"`
	DigitalZoomRatio            string `xml:"digital_zoom_ratio"`
	SubsecTime                  string `xml:"subsec_time"`
	JpegInterchangeFormat       string `xml:"jpeg_interchange_format"`
	JpegInterchangeFormatLength string `xml:"jpeg_interchange_format_length"`
	Compression                 string `xml:"compression"`
}

// URL contains the url of the image and it's type as read from XML
type URL struct {
	URL  string `xml:",innerxml"`
	Type string `xml:"type,attr"`
}
