$(call add_json_map, Pixeldust)
add_json_val_default = $(call add_json_val, $(1), $(if $(strip $(2)), $(2), $(3)))

$(call add_json_str,  Target_shim_libs,                      $(TARGET_LD_SHIM_LIBS))
$(call add_json_bool, Target_uses_color_metadata,            $(filter true,$(TARGET_USES_COLOR_METADATA)))
$(call add_json_val_default, Bootloader_message_offset,      $(BOOTLOADER_MESSAGE_OFFSET), 0)
$(call end_json_map)
