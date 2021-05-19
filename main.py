import os
import msvcrt


def lock_file(f):
    msvcrt.locking(f.fileno(), msvcrt.LK_RLCK, 1)


def unlock_file(f):
    msvcrt.locking(f.fileno(), msvcrt.LK_UNLCK, 1)


def work_with_file(name):
    fd = os.open('foo.txt', os.O_RDWR|os.O_CREAT)
    os.close(fd)

