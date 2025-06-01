{
  description = "Melodex";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = {nixpkgs, ...}: let
    supportedSystems = ["x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin"];

    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

    nixpkgsFor = forAllSystems (system: import nixpkgs {inherit system;});
  in {
    packages = forAllSystems (system: let
      pkgs = nixpkgsFor.${system};
    in rec {
      melodex = pkgs.buildGoModule {
        pname = "Melodex";
        version = "0.0.8";
        src = ./.;
        # vendorHash = nixpkgs.lib.fakeHash;
        vendorHash = "sha256-wztHl8V0xQbPUVPIoiyqNEuQt0TxoOJVQyYYMrW4RhU=";
        buildInputs = [
          pkgs.libvlc
        ];
      };

      default = melodex;
    });
  };
}
