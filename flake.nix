{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    crane = {
        url = "github:ipetkov/crane";
        inputs = {
            nixpkgs.follows = "nixpkgs";
        };
    };
    rust-overlay = {
      url = "github:oxalica/rust-overlay";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
  };

  outputs = { self, nixpkgs, flake-utils, rust-overlay, crane }:
    flake-utils.lib.eachDefaultSystem
    (system:
        let
            overlays = [(import rust-overlay)];
            pkgs = import nixpkgs {
                inherit system overlays;
            };
            rustToolchain = pkgs.rust-bin.nightly.latest.default;
            craneLib = (crane.mkLib pkgs).overrideToolchain rustToolchain;
            my-crate = craneLib.buildPackage {
                src = craneLib.cleanCargoSource (craneLib.path ./.);
                strictDeps = true;

                buildInputs = [
                    # Add additional build inputs here
                    pkgs.openssl pkgs.sqlite
                ] ++ pkgs.lib.optionals pkgs.stdenv.isDarwin [
                    # Additional darwin specific inputs can be set here
                    pkgs.libiconv
                ];
            };
        in
        with pkgs;
        {
            checks = {
                inherit my-crate;
            };
            packages = {
                default = my-crate;
            };
            devShells.default = mkShell {
                buildInputs = with pkgs; [rustToolchain openssl sqlite];
            };
        }
    );
}
