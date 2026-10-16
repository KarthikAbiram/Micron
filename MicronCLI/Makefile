# -*- coding: utf-8 -*-
# Build MicronCLI and package its skill file.
# Run from the repository root with: make -C MicronCLI build

.PHONY: build clean

build:
	@powershell -NoProfile -Command "$$ErrorActionPreference = 'Stop'; $$root = $$PWD; $$sourceDir = Join-Path $$root 'Source'; $$buildRoot = Join-Path $$root '..\Builds\MicronCLI\EXE'; $$binDir = Join-Path $$buildRoot 'bin'; $$skillDir = Join-Path $$buildRoot 'skills\microncli'; $$binary = Join-Path $$binDir 'micronCLI.exe'; $$skillSource = Join-Path $$root 'Docs\SKILL.md'; $$skillOutput = Join-Path $$skillDir 'SKILL.md'; New-Item -ItemType Directory -Force -Path $$binDir, $$skillDir | Out-Null; Push-Location $$sourceDir; try { & go build -o $$binary main.go } finally { Pop-Location }; Copy-Item $$skillSource $$skillOutput -Force"

clean:
	@powershell -NoProfile -Command "$$ErrorActionPreference = 'Stop'; $$buildRoot = Join-Path $$PWD '..\Builds\MicronCLI\EXE'; if (Test-Path $$buildRoot) { Remove-Item $$buildRoot -Recurse -Force }"
