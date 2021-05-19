import msvcrt

def lock_file(f):
    msvcrt.locking(f.fileno(), msvcrt.LK_RLCK, 1)


def unlock_file(f):
    msvcrt.locking(f.fileno(), msvcrt.LK_UNLCK, 1)
