terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_file_folder" "my_file_parent_folder" {
  name = "My Parent Folder"
}

resource "splight_file_folder" "my_file_folder" {
  name   = "My Folder"
  parent = splight_file_folder.my_file_parent_folder.id
}
