# Build the plugin
make build

chmod +x packer-plugin-syft
# Create the directory structure that Packer expects
mkdir -p ~/.packer.d/plugins/github.com/gruffwizard/syft/1.0.0/darwin_arm64/

# Copy the plugin to the right location
cp packer-plugin-syft ~/.packer.d/plugins/github.com/gruffwizard/syft/1.0.0/darwin_arm64/
cp packer-plugin-syft ~/.packer.d/plugins/
# run test

