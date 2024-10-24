#!/usr/bin/env python3

import os

def process_go_file(file_path):
    with open(file_path, 'r') as file:
        lines = file.readlines()

    # Check the first five lines for the SPDX identifier
    spdx_present = any(line.startswith("// SPDX-License-Identifier:") for line in lines[:5])

    if not spdx_present:
        # Add the SPDX identifier at the beginning
        new_lines = ["// SPDX-License-Identifier: GPL-3.0-or-later\n", "\n"] + lines
        with open(file_path, 'w') as file:
            file.writelines(new_lines)
        print(f"Added SPDX identifier to {file_path}")
    else:
        print(f"SPDX identifier already present in {file_path}")

def walk_directory_and_process_files(directory):
    for root, _, files in os.walk(directory):
        for file in files:
            if file.endswith('.go'):
                file_path = os.path.join(root, file)
                process_go_file(file_path)

if __name__ == "__main__":
    current_directory = os.getcwd()
    walk_directory_and_process_files(current_directory)
