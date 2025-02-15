module.exports = {
  apps: [
    {
      name: "open-api-golang",
      script: "./open_api_golang",
      cwd: "/var/www/open_api_golang",
      autorestart: true,
      watch: false,
      env_file: "/var/www/open_api_golang/.env",
      env: {
        PATH: "/usr/local/go/bin:/usr/bin:/bin"
      },
      pre_start: "go mod tidy && go build -o ./open_api_golang -buildvcs=false"
    }
  ]
};
