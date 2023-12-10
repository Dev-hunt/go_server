# go_server

step 1: Install go with dependecied
Step 2: Create a file of webserver.go
step 3: code is written using error handling, logs and every code interruption
step 4: run the code with  "go run webserve.go"
step 5: Server will start running at 8000, now move frwd to check it though postman or cURL

link according to server deployment  

curl --location 'http://localhost:8000/process-concurrent' \
--header 'Content-Type: application/json' \
--header 'Cookie: .Tunnels.Relay.WebForwarding.Cookies=CfDJ8M67rYfw57hCj5sJjtQyecG_2RZmB2zO4GN8kUzRujryZ1WNIZy3zQgEC-Hy3dzhOssx1Ma967ujSgeICQpLnOBEDD-iBrrWDpGmkHvLNeZIBTA0IFZ5Z2m9RS-WuNfkbQdSlMr2OMo2jtCM77ToVE83rnNxdT8BYmcBcg4KdvbS2v8OX72EZFzTwtxzoPTV9VAEGjKsbhxzDDL46QPL-0Y3ljyPEKC2zTlb4estq7swdV24lALCVik4mQhE1xjgyEIfUoVnCf9nYFduQaU0efw9ri1Q6WVHWyIxl9Sm3BpIf_U-8IdI_Hm6p4LENEAnmk1AVJBOF4nUWHtQsm8DOSM5gaYezPoR7hKWzf2dFbidbcUPZ9Dq9gJh6_U1wfxzl0XFza54qJrHguCUnfhGjKw-OgwbjeNWGGBtdFcvBt29HyhypD8fnFUMGjSqDvIoG7Xk7J6EOgmGxgHCX_sOH2S9Le-GbYoGDPbJdF3IIoZyAZR6bbFxve48Que-LfL_r7HqFPorvq-mO412RBzyI1UroM4OWnWCefkmaSFSGNNTb2D1oQJNDHPuHMc9o8p3KED-U0bCBRTSTCsAYafTFktYG5lH86_H2uCLIzJkW2wV1JjwzvrsKPD0_vauJdHC32J8qokPlAZKkCV8usQHmzIqhjm1p6vaA89fwa_Px0hNHwoz88qlSYRpdoYZX3ke21QrNVpWZVnvFjaKlgtRNDURoLsD6gweK4N6EfPNYj7oP339DJemNhB7PXzS31vQ15fxB8TXBqSHUrdIQC1avhamgk25okxQaYOtO1zl6U-5BqS0zspDblbtWBHpkXrMCObeSggaCFb_AoRyuz7A5u1elXXJywM-yr0mGnFZPl_cukdHxegXRzxUDRCM3ulnB_l2U2GnXghsbthBUijluInMCBVvELdbTS0wizuoch1yNv3rgEtaU1uylc2Y8PDAdB947wlQ-aWBP97wbj69T7UTgGPQIgEN-wYpkquEm-ik' \
--data '{
  "to_sort": [[9, 2, 3], [4, 8, 6, 45], [7, 22, 9, 67, 98]]
}'


curl --location 'http://localhost:8000/process-single' \
--header 'Content-Type: application/json' \
--header 'Cookie: .Tunnels.Relay.WebForwarding.Cookies=CfDJ8M67rYfw57hCj5sJjtQyecG_2RZmB2zO4GN8kUzRujryZ1WNIZy3zQgEC-Hy3dzhOssx1Ma967ujSgeICQpLnOBEDD-iBrrWDpGmkHvLNeZIBTA0IFZ5Z2m9RS-WuNfkbQdSlMr2OMo2jtCM77ToVE83rnNxdT8BYmcBcg4KdvbS2v8OX72EZFzTwtxzoPTV9VAEGjKsbhxzDDL46QPL-0Y3ljyPEKC2zTlb4estq7swdV24lALCVik4mQhE1xjgyEIfUoVnCf9nYFduQaU0efw9ri1Q6WVHWyIxl9Sm3BpIf_U-8IdI_Hm6p4LENEAnmk1AVJBOF4nUWHtQsm8DOSM5gaYezPoR7hKWzf2dFbidbcUPZ9Dq9gJh6_U1wfxzl0XFza54qJrHguCUnfhGjKw-OgwbjeNWGGBtdFcvBt29HyhypD8fnFUMGjSqDvIoG7Xk7J6EOgmGxgHCX_sOH2S9Le-GbYoGDPbJdF3IIoZyAZR6bbFxve48Que-LfL_r7HqFPorvq-mO412RBzyI1UroM4OWnWCefkmaSFSGNNTb2D1oQJNDHPuHMc9o8p3KED-U0bCBRTSTCsAYafTFktYG5lH86_H2uCLIzJkW2wV1JjwzvrsKPD0_vauJdHC32J8qokPlAZKkCV8usQHmzIqhjm1p6vaA89fwa_Px0hNHwoz88qlSYRpdoYZX3ke21QrNVpWZVnvFjaKlgtRNDURoLsD6gweK4N6EfPNYj7oP339DJemNhB7PXzS31vQ15fxB8TXBqSHUrdIQC1avhamgk25okxQaYOtO1zl6U-5BqS0zspDblbtWBHpkXrMCObeSggaCFb_AoRyuz7A5u1elXXJywM-yr0mGnFZPl_cukdHxegXRzxUDRCM3ulnB_l2U2GnXghsbthBUijluInMCBVvELdbTS0wizuoch1yNv3rgEtaU1uylc2Y8PDAdB947wlQ-aWBP97wbj69T7UTgGPQIgEN-wYpkquEm-ik' \
--data '{
  "to_sort": [[9, 2, 3], [4, 8, 6, 45], [7, 22, 9, 67, 98]]
}'


Now move forward, make a go init by using go mod init "filename"

step 6: make a docer file with all the requirement and cmds mentiond
step 7: run "docker build -t go-server-image ."  to make a image of docker

step 8: login docker docker login -u <your-docker-username> -p <your-new-access-token>
step 9: make a repo in docker and push  the image with tag docker push devansh19/go_server:v1.0

step 10: check that docker image runnig or not by "docker run -p 8000:8000 go-server-image"

step 11: now open render, login then make a webservice. pull your latest version of docker image, then publish it and configure your domain.
