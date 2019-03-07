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

# IFS= preserve newlines
IFS= manifest=$(curl -s https://api.github.com/repos/jgramoll/terraform-provider-spinnaker/releases/latest)

url=$(echo $manifest \
| grep "browser_download_url.*${kernel}_${arch}" \
| cut -d '"' -f 4 \
)
version=$(echo $manifest \
| grep tag_name \
| cut -d '"' -f 4 \
)

if [ -z ${url} ]; then
  echo "no build for this kernel/arch: ${kernel}_${arch}"
  exit 1
fi

dest_file="terraform-provider-spinnaker_$version"
curl $url -L -o $dest_file
chmod +x $dest_file

mkdir -p $terraform_plugins
mv $dest_file $terraform_plugins/