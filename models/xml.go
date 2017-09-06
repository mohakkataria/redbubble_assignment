// Package model contains all the necessary type declarations required for the Domain objects of the assignment
package models

import "encoding/xml"

// XML is the root level of XML to be read.
type XML struct {
	Works xml.Name `xml:"works>work"`
	Work []Work `xml:"work"`
}

// Work is definition of single piece of work/image
type Work struct {
	Id int `xml:"id"`
	Filename string `xml:"filename"`
	ImageWidth int `xml:"image_width"`
	ImageHeight int `xml:"image_height"`
	URLs []URL `xml:"urls>url"`
	EXIF EXIF `xml:"exif"`
}

// EXIF contains all the metadata of the image taken
type EXIF struct {
	Flash string `xml:"flash"`
	Sharpness string `xml:"sharpness"`
	Pixel_y_dimension string `xml:"pixel_y_dimension"`
	Date_time_digitized string `xml:"date_time_digitized"`
	Exposure_bias_value string `xml:"exposure_bias_value"`
	Focal_length_in_35mm_film string `xml:"focal_length_in_35mm_film"`
	Subsec_time_orginal string `xml:"subsec_time_orginal"`
	Model string `xml:"model"`
	Software string `xml:"software"`
	Focal_length string `xml:"focal_length"`
	Max_aperture_value string `xml:"max_aperture_value"`
	Scene_capture_type string `xml:"scene_capture_type"`
	Subsec_time_digitized string `xml:"subsec_time_digitized"`
	Iso_speed_ratings string `xml:"iso_speed_ratings"`
	Date_time string `xml:"date_time"`
	Compressed_bits_per_pixel string `xml:"compressed_bits_per_pixel"`
	Subject_distance_range string `xml:"subject_distance_range"`
	Custom_rendered string `xml:"custom_rendered"`
	Make string `xml:"make"`
	Gain_control string `xml:"gain_control"`
	Orientation string `xml:"orientation"`
	Exposure_time string `xml:"exposure_time"`
	Exposure_mode string `xml:"exposure_mode"`
	X_resolution string `xml:"x_resolution"`
	Metering_mode string `xml:"metering_mode"`
	Sensing_method string `xml:"sensing_method"`
	Contrast string `xml:"contrast"`
	Color_space string `xml:"color_space"`
	F_number string `xml:"f_number"`
	White_balance string `xml:"white_balance"`
	User_comment string `xml:"user_comment"`
	Y_resolution string `xml:"y_resolution"`
	Resolution_unit string `xml:"resolution_unit"`
	Light_source string `xml:"light_source"`
	Saturation string `xml:"saturation"`
	Pixel_x_dimension string `xml:"pixel_x_dimension"`
	Date_time_original string `xml:"date_time_original"`
	Ycb_cr_positioning string `xml:"ycb_cr_positioning"`
	Exposure_program string `xml:"exposure_program"`
	Digital_zoom_ratio string `xml:"digital_zoom_ratio"`
	Subsec_time string `xml:"subsec_time"`
	Jpeg_interchange_format string `xml:"jpeg_interchange_format"`
	Jpeg_interchange_format_length string `xml:"jpeg_interchange_format_length"`
	Compression string `xml:"compression"`
}

// URL contains the url of the image and it's type as read from XML
type URL struct {
	URL string `xml:",innerxml"`
	Type string `xml:"type,attr"`
}