{
    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = { self, nixpkgs, flake-utils }: 
        flake-utils.lib.eachDefaultSystem (system:
            let
              name = "simple-go-project";
              src = ./.;
              pkgs = nixpkgs.legacyPackages.${system};
            in 
            {
                packages.default = pkgs.mkShell
                {
                    nativeBuildInputs = with pkgs; [ 
                        go 
                        gnumake
                    ];
                };
            }
        );
}
