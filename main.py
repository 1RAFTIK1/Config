from zipfile import ZipFile
import sys

def delete_symbol(path):
    for letter in path:
        if letter == "/":
            path = path[1:]
        else:
            break
        return path

def ls (path, files):
    path = delete_symbol(path)
    for file in files:
        if path in file.filename:
            file_names = file.filename[len(path):].split("/")
            file_names = list(filter(None, file_names))
            if len(file_names) > 1 or not file_names:
                continue
            print("\033[33m{}\033[0m".format(file_names[0]))

