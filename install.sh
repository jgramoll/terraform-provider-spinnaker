terraform_plugins=~/.terraform.d/plugins/darwin_amd64/

arch=$(uname -m)
if [ $arch != "x86_64" ] && [ $arch != "i386" ]; then
  echo "no build for this architecture: $arch"
  exit 1
fi

kernel=$(uname -s)
if [ $kernel != "Darwin" ] && [ $kernel != "Linux" ]; then
  echo "no build for this kernel: $kernel"
  exit 1
fi

manifest=$(curl -s https://api.github.com/repos/jgramoll/terraform-provider-spinnaker/releases/latest)
version=$(echo $manifest | jq --raw-output ".name")
url=$(echo $manifest | jq --raw-output ".assets[] | select(.name | contains(\"${kernel}_${arch}\")) | .browser_download_url")

if [ -z ${url} ]; then
  echo "no build for this kernel/arch: ${kernel}_${arch}"
  exit 1
fi

dest_file="terraform-provider-spinnaker_$version"
# curl $url -L -o $dest_file

mkdir -p $terraform_plugins
mv $dest_file $terraform_plugins/
