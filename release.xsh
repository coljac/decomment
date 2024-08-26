#!env xonsh
import os

tag = $(git tag | tail -1)

for d in $(ls).split():
    if os.path.isdir(d):
        cd @(d)
        for arch in $(ls).split():
            if os.path.isdir(arch):
                cd @(arch)
                zip ../../@(f"dec-{d}-{arch}-{tag}.zip") *
                cd ..
        cd ..
mv *zip ~/Downloads/

