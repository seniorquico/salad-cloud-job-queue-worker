# Mandelbrot Workload Sample

This directory contains examples of different strategies that may be used to embed and run the SaladCloud Job Queue Worker in an existing workload container image.

The `base` directory contains a sample workload, a simple Mandelbrot image generator that runs as a HTTP service.

The `with-s6-overlay` directory contains an example integration strategy that downloads and installs the SaladCloud Job Queue Worker binary from the GitHub release and uses [s6-overlay](https://github.com/just-containers/s6-overlay) to run both the HTTP service and the SaladCloud Job Queue Worker.

The `with-shell-script` directory contains an example integration strategy that downloads and installs the SaladCloud Job Queue Worker binary from the GitHub release and uses a simple shell script to run both the HTTP service and the SaladCloud Job Queue Worker.

## Development

The following prerequisites are required:

- [Git](https://git-scm.com/downloads)
- [Docker](https://www.docker.com/get-started/)

1. Clone the repository.

   ```sh
   git clone https://github.com/SaladTechnologies/salad-cloud-job-queue-worker.git
   ```

2. Build the workload container image.

   ```sh
   cd samples/mandelbrot/base
   docker image build -t mandelbrot:latest .
   ```

3. Build the example image.

   ```sh
   cd samples/mandelbrot/with-s6-overlay
   docker image build -t mandelbrot-worker:latest .
   ```

   Alternatively, change to the `with-shell-script` directory before building the example image.
