package android
type Product_variables struct {
	Target_shim_libs struct {
		Cppflags []string
	}
	Target_uses_color_metadata struct {
		Cppflags []string
	}
	Uses_qcom_um_family struct {
		Cflags []string
		Srcs []string
	}
	Uses_qcom_um_3_18_family struct {
		Include_dirs []string
		Header_libs []string
		Shared_libs []string
	}
	Uses_qcom_um_4_4_family struct {
		Include_dirs []string
		Header_libs []string
		Shared_libs []string
	}
	Uses_qcom_um_4_9_family struct {
		Include_dirs []string
		Header_libs []string
		Shared_libs []string
	}
	Uses_qcom_um_4_14_family struct {
		Include_dirs []string
		Header_libs []string
		Shared_libs []string
	}
}

type ProductVariables struct {
	Target_shim_libs					*string `json:",omitempty"`
	Target_uses_color_metadata				*bool `json:",omitempty"`
	Uses_qcom_um_family					*bool `json:",omitempty"`
	Uses_qcom_um_3_18_family				*bool `json:",omitempty"`
	Uses_qcom_um_4_4_family					*bool `json:",omitempty"`
	Uses_qcom_um_4_9_family					*bool `json:",omitempty"`
	Uses_qcom_um_4_14_family				*bool `json:",omitempty"`
}
