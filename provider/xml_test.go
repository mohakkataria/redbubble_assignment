package provider

import (
	"testing"
	"fmt"
)

func TestXML_ConvertToImages(t *testing.T) {
	xmlProvider := XML{XML:[]byte(`<works>
		<work>
		<id>31820</id>
		<filename>162042.jpg</filename>
		<image_width>2206</image_width>
		<image_height>2205</image_height>
		<urls>
		<url type="small">
		http://ih1.redbubble.net/work.31820.1.flat,135x135,075,f.jpg
		</url>
		<url type="medium">
		http://ih1.redbubble.net/work.31820.1.flat,300x300,075,f.jpg
		</url>
		<url type="large">
		http://ih1.redbubble.net/work.31820.1.flat,550x550,075,f.jpg
		</url>
		</urls>
		<exif>
		<flash>0</flash>
		<sharpness>2</sharpness>
		<pixel_y_dimension>2205</pixel_y_dimension>
		<date_time_digitized>Mon Jun 04 14:33:59 +1000 2007</date_time_digitized>
		<exposure_bias_value>0</exposure_bias_value>
		<focal_length_in_35mm_film>93</focal_length_in_35mm_film>
		<subsec_time_orginal>00</subsec_time_orginal>
		<model>NIKON D80</model>
		<software>Adobe Photoshop CS3 Macintosh</software>
		<focal_length>62</focal_length>
		<max_aperture_value>47/10</max_aperture_value>
		<scene_capture_type>0</scene_capture_type>
		<subsec_time_digitized>00</subsec_time_digitized>
		<iso_speed_ratings>100</iso_speed_ratings>
		<date_time>Sun Aug 26 12:55:02 +1000 2007</date_time>
		<compressed_bits_per_pixel>4</compressed_bits_per_pixel>
		<subject_distance_range>0</subject_distance_range>
		<custom_rendered>0</custom_rendered>
		<make>NIKON CORPORATION</make>
		<gain_control>0</gain_control>
		<orientation>1</orientation>
		<exposure_time>1/10</exposure_time>
		<exposure_mode>0</exposure_mode>
		<x_resolution>300</x_resolution>
		<metering_mode>5</metering_mode>
		<sensing_method>2</sensing_method>
		<contrast>0</contrast>
		<color_space>1</color_space>
		<f_number>5</f_number>
		<white_balance>0</white_balance>
		<user_comment/>
		<y_resolution>300</y_resolution>
		<resolution_unit>2</resolution_unit>
		<light_source>0</light_source>
		<saturation>2</saturation>
		<pixel_x_dimension>2206</pixel_x_dimension>
		<date_time_original>Mon Jun 04 14:33:59 +1000 2007</date_time_original>
		<ycb_cr_positioning>2</ycb_cr_positioning>
		<exposure_program>4</exposure_program>
		<digital_zoom_ratio>1</digital_zoom_ratio>
		<subsec_time>00</subsec_time>
		<jpeg_interchange_format>978</jpeg_interchange_format>
		<jpeg_interchange_format_length>7171</jpeg_interchange_format_length>
		<x_resolution>72</x_resolution>
		<compression>6</compression>
		<y_resolution>72</y_resolution>
		<resolution_unit>2</resolution_unit>
		</exif>
		</work></works>`)}
	actual := xmlProvider.ConvertToImages()
	fmt.Println(actual)
	if len(actual) == 1 {
		if actual[0].Make != "NIKON CORPORATION" {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", "NIKON CORPORATION", actual[0].Make)
		}

		if actual[0].Model != "NIKON D80" {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", "NIKON D80", actual[0].Model)
		}

		if actual[0].Thumbnail != "http://ih1.redbubble.net/work.31820.1.flat,135x135,075,f.jpg" {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", "http://ih1.redbubble.net/work.31820.1.flat,135x135,075,f.jpg", actual[0].Thumbnail)
		}

		if actual[0].ID != 31820 {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", 31820, actual[0].ID)
		}
	} else {

	}
}

