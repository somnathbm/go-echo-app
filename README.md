# a simple Go Echo server app
This is a test app. This app to be CI-ed to DockerHub by Github action.

Then from GitHub Webhook (will be hosted on another repo), ArgoCD can take over and deploy to a EKS cluster.
Then eventually Argo Rollouts will supervise the susequent rollouts