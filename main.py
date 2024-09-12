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
def cd(path, extension_path, files):
    if "root:" in extension_path:
        path = extension_path[len("root:")]
    else:
        path += "/" + extension_path
    path = delete_symbol(path)

    global local_path

    if path == "":
        local_path = ""
        return True

    if "../" in path:
        local_path = local_path[:len(local_path) - len(local_path.split("/")[-1]) - 1]
        return True

    for file in files:
        if path in file.filename:
            local_path = "/" + path
            return True
    return False


def cat(path, extension_path, zip_file):
    if "root:" in extension_path:
        path = extension_path[len("root:")]
    else:
        path += "/" + extension_path
    path = delete_symbol(path)

    flag = False
    for file in ZipFile(zip_file).filelist:
        if path in file.filename:
            flag = True
            with ZipFile(zip_file) as files:
                with files.open(path,'r') as file:
                    for line in file.readlines():
                        print(line.decode('utf8').strip())
    if not flag:
        print("\033[33m{}\033[0m".format("Can`t open this file"))
