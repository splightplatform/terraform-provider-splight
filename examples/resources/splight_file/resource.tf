resource "spl_file" "FileInnerTest" {
  file        = "./variables.tf"
  description = "Sample file for testing inner file"
  parent      = "1234-1234-1234-1234"
}
