{
  description = "gjq - GameJanitor Query Library & CLI";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.buildGoModule {
            pname = "gjq";
            version = "0.1.0";
            src = ./.;
            vendorHash = "sha256-KPSU7vSNvGS+TJgRuCStur0AbaRph4OS5Z50qeCJYAQ=";
            subPackages = [ "cmd/gjq" ];

            meta = {
              description = "GameJanitor Query CLI";
              mainProgram = "gjq";
            };
          };
        });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};

          update-vendor-hash = pkgs.writeShellScriptBin "update-vendor-hash" ''
            go mod vendor
            HASH=$(nix hash path --type sha256 vendor/)
            sed -i "s|vendorHash = \".*\"|vendorHash = \"$HASH\"|" flake.nix
            echo "Updated vendorHash to $HASH"
          '';
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gopls
              gotools
              update-vendor-hash
            ];

            shellHook = ''
              gjq() { go run ./cmd/gjq "$@"; }
              export -f gjq
            '';
          };
        });
    };
}
