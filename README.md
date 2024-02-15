# Steps to Reproduce

1. Use the EdgeX Foundry Napa release [no-security compose file](https://github.com/edgexfoundry/edgex-compose/blob/napa/docker-compose-no-secty.yml) (copied here)

   ```sh
   docker compose -f docker-compose-no-secty.yml up -d
   ```

2. Run the app service in this repository

   ```sh
   EDGEX_SECURITY_SECRET_STORE=false go run . -cp
   ```

3. Check the logs

   ```txt
   level=INFO ts=2024-02-15T21:53:03.825794909Z app=service-key source=service.go:587 msg="Service started in: 7.997465849s"
   level=INFO ts=2024-02-15T21:53:03.825799409Z app=service-key source=configupdates.go:48 msg="Waiting for App Service configuration updates..."
   level=INFO ts=2024-02-15T21:53:03.82764621Z app=service-key source=main.go:25 msg="Event Count: 17"
   level=ERROR ts=2024-02-15T21:53:03.82769241Z app=service-key source=main.go:29 msg=">>>>>>>>>> reading client is nil <<<<<<<<<<"
   ```
