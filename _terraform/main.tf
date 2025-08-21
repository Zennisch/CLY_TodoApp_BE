resource "google_compute_instance" "jec_vm" {
  name         = "jec"
  machine_type = "e2-medium"
  zone         = "${var.gcp_region}-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-2404-lts-amd64"
      size  = 10
      type  = "pd-standard"
    }
  }

  network_interface {
    network = "default"
    access_config {}
  }

  metadata = {
    ssh-keys = var.ssh_public_key
  }

  tags = ["jec", "allow-ssh", "allow-app-ports"]
}

resource "google_compute_firewall" "allow_ssh" {
  name    = "allow-ssh-jec"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["allow-ssh"]
}

resource "google_compute_firewall" "allow_app_ports_8000" {
  name    = "allow-app-ports-8000-jec"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["8000"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["allow-app-ports"]
}

resource "google_compute_firewall" "allow_app_ports_443" {
  name    = "allow-app-ports-443-jec"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["443"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["allow-app-ports"]
}

resource "google_compute_firewall" "allow_app_ports_80" {
  name    = "allow-app-ports-80-jec"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["allow-app-ports"]
}