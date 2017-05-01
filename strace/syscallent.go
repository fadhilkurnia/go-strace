package strace

const (
	TRACE_FILE                  = 001
	TRACE_IPC                   = 002
	TRACE_NETWORK               = 004
	TRACE_PROCESS               = 010
	TRACE_SIGNAL                = 020
	TRACE_DESC                  = 040
	TRACE_MEMORY                = 0100
	SYSCALL_NEVER_FAILS         = 0200
	STACKTRACE_INVALIDATE_CACHE = 0400
	STACKTRACE_CAPTURE_ON_ENTER = 01000
	TRACE_INDIRECT_SUBCALL      = 02000
	COMPAT_SYSCALL_TYPES        = 04000
	TRACE_SCHED                 = 010000

	MAX_ARGS = 6

	TD  = TRACE_DESC
	TF  = TRACE_FILE
	TI  = TRACE_IPC
	TN  = TRACE_NETWORK
	TP  = TRACE_PROCESS
	TS  = TRACE_SIGNAL
	TM  = TRACE_MEMORY
	TSC = TRACE_SCHED
	NF  = SYSCALL_NEVER_FAILS
	MA  = MAX_ARGS
	SI  = STACKTRACE_INVALIDATE_CACHE
	SE  = STACKTRACE_CAPTURE_ON_ENTER
	CST = COMPAT_SYSCALL_TYPES

	ARG_INT = 1
	ARG_PTR = 2
	ARG_STR = 3
)

type sysent struct {
	Nargs    uint
	SysFlags int
	SysName  string
	ArgTypes []int
}

var syscallent []sysent = []sysent{
	0:   sysent{3, TD, "read", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	1:   sysent{3, TD, "write", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	2:   sysent{3, TD | TF, "open", []int{ARG_STR, ARG_INT, ARG_INT, -1, -1, -1}},
	3:   sysent{1, TD, "close", []int{ARG_INT, -1, -1, -1, -1, -1}},
	4:   sysent{2, TF, "stat", []int{ARG_STR, ARG_PTR, -1, -1, -1, -1}},
	5:   sysent{2, TD, "fstat", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	6:   sysent{2, TF, "lstat", []int{ARG_STR, ARG_PTR, -1, -1, -1, -1}},
	7:   sysent{3, TD, "poll", []int{ARG_PTR, ARG_INT, ARG_INT, -1, -1, -1}},
	8:   sysent{3, TD, "lseek", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	9:   sysent{6, TD | TM | SI, "mmap", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT}},
	10:  sysent{3, TM | SI, "mprotect", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	11:  sysent{2, TM | SI, "munmap", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	12:  sysent{1, TM | SI, "brk", []int{ARG_INT, -1, -1, -1, -1, -1}},
	13:  sysent{4, TS, "rt_sigaction", []int{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT, -1, -1}},
	14:  sysent{4, TS, "rt_sigprocmask", []int{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT, -1, -1}},
	15:  sysent{0, TS, "rt_sigreturn", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	16:  sysent{3, TD, "ioctl", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	17:  sysent{4, TD, "pread64", []int{ARG_INT, ARG_STR, ARG_INT, ARG_INT, -1, -1}},
	18:  sysent{4, TD, "pwrite64", []int{ARG_INT, ARG_STR, ARG_INT, ARG_INT, -1, -1}},
	19:  sysent{3, TD, "readv", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	20:  sysent{3, TD, "writev", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	21:  sysent{2, TF, "access", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	22:  sysent{1, TD, "pipe", []int{ARG_PTR, -1, -1, -1, -1, -1}},
	23:  sysent{5, TD, "select", []int{ARG_INT, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, -1}},
	24:  sysent{0, TSC, "sched_yield", []int{-1, -1, -1, -1, -1, -1}},
	25:  sysent{5, TM | SI, "mremap", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1}},
	26:  sysent{3, TM, "msync", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	27:  sysent{3, TM, "mincore", []int{ARG_INT, ARG_INT, ARG_PTR, -1, -1, -1}},
	28:  sysent{3, TM, "madvise", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	29:  sysent{3, TI, "shmget", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	30:  sysent{3, TI | TM | SI, "shmat", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	31:  sysent{3, TI, "shmctl", []int{ARG_INT, ARG_INT, ARG_PTR, -1, -1, -1}},
	32:  sysent{1, TD, "dup", []int{ARG_INT, -1, -1, -1, -1, -1}},
	33:  sysent{2, TD, "dup2", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	34:  sysent{0, TS, "pause", []int{-1, -1, -1, -1, -1, -1}},
	35:  sysent{2, 0, "nanosleep", []int{ARG_PTR, ARG_PTR, -1, -1, -1, -1}},
	36:  sysent{2, 0, "getitimer", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	37:  sysent{1, 0, "alarm", []int{ARG_INT, -1, -1, -1, -1, -1}},
	38:  sysent{3, 0, "setitimer", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	39:  sysent{0, NF, "getpid", []int{-1, -1, -1, -1, -1, -1}},
	40:  sysent{4, TD | TN, "sendfile", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_INT, -1, -1}},
	41:  sysent{3, TN, "socket", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	42:  sysent{3, TN, "connect", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	43:  sysent{3, TN, "accept", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	44:  sysent{6, TN, "sendto", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, ARG_PTR, ARG_INT}},
	45:  sysent{6, TN, "recvfrom", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}},
	46:  sysent{3, TN, "sendmsg", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	47:  sysent{3, TN, "recvmsg", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	48:  sysent{2, TN, "shutdown", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	49:  sysent{3, TN, "bind", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	50:  sysent{2, TN, "listen", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	51:  sysent{3, TN, "getsockname", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	52:  sysent{3, TN, "getpeername", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	53:  sysent{4, TN, "socketpair", []int{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, -1, -1}},
	54:  sysent{5, TN, "setsockopt", []int{ARG_INT, ARG_INT, ARG_INT, ARG_STR, ARG_INT, -1}},
	55:  sysent{5, TN, "getsockopt", []int{ARG_INT, ARG_INT, ARG_INT, ARG_STR, ARG_PTR, -1}},
	56:  sysent{5, TP, "clone", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, ARG_INT, -1}},
	57:  sysent{0, TP, "fork", []int{-1, -1, -1, -1, -1, -1}},
	58:  sysent{0, TP, "vfork", []int{-1, -1, -1, -1, -1, -1}},
	59:  sysent{3, TF | TP | SE | SI, "execve", []int{ARG_STR, ARG_PTR, ARG_PTR, -1, -1, -1}},
	60:  sysent{1, TP | SE, "exit", []int{ARG_INT, -1, -1, -1, -1, -1}},
	61:  sysent{4, TP, "wait4", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, -1, -1}},
	62:  sysent{2, TS, "kill", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	63:  sysent{1, 0, "uname", []int{ARG_PTR, -1, -1, -1, -1, -1}},
	64:  sysent{3, TI, "semget", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	65:  sysent{3, TI, "semop", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	66:  sysent{4, TI, "semctl", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1, -1}},
	67:  sysent{1, TI | TM | SI, "shmdt", []int{ARG_STR, -1, -1, -1, -1, -1}},
	68:  sysent{2, TI, "msgget", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	69:  sysent{4, TI, "msgsnd", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, -1, -1}},
	70:  sysent{5, TI, "msgrcv", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, ARG_INT, -1}},
	71:  sysent{3, TI, "msgctl", []int{ARG_INT, ARG_INT, ARG_PTR, -1, -1, -1}},
	72:  sysent{3, TD, "fcntl", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	73:  sysent{2, TD, "flock", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	74:  sysent{1, TD, "fsync", []int{ARG_INT, -1, -1, -1, -1, -1}},
	75:  sysent{1, TD, "fdatasync", []int{ARG_INT, -1, -1, -1, -1, -1}},
	76:  sysent{2, TF, "truncate", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	77:  sysent{2, TD, "ftruncate", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	78:  sysent{3, TD, "getdents", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	79:  sysent{2, TF, "getcwd", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	80:  sysent{1, TF, "chdir", []int{ARG_STR, -1, -1, -1, -1, -1}},
	81:  sysent{1, TD, "fchdir", []int{ARG_INT, -1, -1, -1, -1, -1}},
	82:  sysent{2, TF, "rename", []int{ARG_STR, ARG_STR, -1, -1, -1, -1}},
	83:  sysent{2, TF, "mkdir", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	84:  sysent{1, TF, "rmdir", []int{ARG_STR, -1, -1, -1, -1, -1}},
	85:  sysent{2, TD | TF, "creat", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	86:  sysent{2, TF, "link", []int{ARG_STR, ARG_STR, -1, -1, -1, -1}},
	87:  sysent{1, TF, "unlink", []int{ARG_STR, -1, -1, -1, -1, -1}},
	88:  sysent{2, TF, "symlink", []int{ARG_STR, ARG_STR, -1, -1, -1, -1}},
	89:  sysent{3, TF, "readlink", []int{ARG_STR, ARG_STR, ARG_INT, -1, -1, -1}},
	90:  sysent{2, TF, "chmod", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	91:  sysent{2, TD, "fchmod", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	92:  sysent{3, TF, "chown", []int{ARG_STR, ARG_INT, ARG_INT, -1, -1, -1}},
	93:  sysent{3, TD, "fchown", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	94:  sysent{3, TF, "lchown", []int{ARG_STR, ARG_INT, ARG_INT, -1, -1, -1}},
	95:  sysent{1, NF, "umask", []int{ARG_INT, -1, -1, -1, -1, -1}},
	96:  sysent{2, 0, "gettimeofday", []int{ARG_PTR, ARG_PTR, -1, -1, -1, -1}},
	97:  sysent{2, 0, "getrlimit", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	98:  sysent{2, 0, "getrusage", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	99:  sysent{1, 0, "sysinfo", []int{ARG_PTR, -1, -1, -1, -1, -1}},
	100: sysent{1, 0, "times", []int{ARG_PTR, -1, -1, -1, -1, -1}},
	101: sysent{4, 0, "ptrace", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1, -1}},
	102: sysent{0, NF, "getuid", []int{-1, -1, -1, -1, -1, -1}},
	103: sysent{3, 0, "syslog", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	104: sysent{0, NF, "getgid", []int{-1, -1, -1, -1, -1, -1}},
	105: sysent{1, 0, "setuid", []int{ARG_INT, -1, -1, -1, -1, -1}},
	106: sysent{1, 0, "setgid", []int{ARG_INT, -1, -1, -1, -1, -1}},
	107: sysent{0, NF, "geteuid", []int{-1, -1, -1, -1, -1, -1}},
	108: sysent{0, NF, "getegid", []int{-1, -1, -1, -1, -1, -1}},
	109: sysent{2, 0, "setpgid", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	110: sysent{0, NF, "getppid", []int{-1, -1, -1, -1, -1, -1}},
	111: sysent{0, 0, "getpgrp", []int{-1, -1, -1, -1, -1, -1}},
	112: sysent{0, 0, "setsid", []int{-1, -1, -1, -1, -1, -1}},
	113: sysent{2, 0, "setreuid", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	114: sysent{2, 0, "setregid", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	115: sysent{2, 0, "getgroups", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	116: sysent{2, 0, "setgroups", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	117: sysent{3, 0, "setresuid", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	118: sysent{3, 0, "getresuid", []int{ARG_PTR, ARG_PTR, ARG_PTR, -1, -1, -1}},
	119: sysent{3, 0, "setresgid", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	120: sysent{3, 0, "getresgid", []int{ARG_PTR, ARG_PTR, ARG_PTR, -1, -1, -1}},
	121: sysent{1, 0, "getpgid", []int{ARG_INT, -1, -1, -1, -1, -1}},
	122: sysent{1, NF, "setfsuid", []int{ARG_INT, -1, -1, -1, -1, -1}},
	123: sysent{1, NF, "setfsgid", []int{ARG_INT, -1, -1, -1, -1, -1}},
	124: sysent{1, 0, "getsid", []int{ARG_INT, -1, -1, -1, -1, -1}},
	125: sysent{2, 0, "capget", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	126: sysent{2, 0, "capset", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	127: sysent{2, TS, "rt_sigpending", []int{ARG_PTR, ARG_INT, -1, -1, -1, -1}},
	128: sysent{4, TS, "rt_sigtimedwait", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_INT, -1, -1}},
	129: sysent{3, TS, "rt_sigqueueinfo", []int{ARG_INT, ARG_INT, ARG_PTR, -1, -1, -1}},
	130: sysent{2, TS, "rt_sigsuspend", []int{ARG_PTR, ARG_INT, -1, -1, -1, -1}},
	131: sysent{2, TS, "sigaltstack", []int{ARG_PTR, ARG_PTR, -1, -1, -1, -1}},
	132: sysent{2, TF, "utime", []int{ARG_STR, ARG_PTR, -1, -1, -1, -1}},
	133: sysent{3, TF, "mknod", []int{ARG_STR, ARG_INT, ARG_INT, -1, -1, -1}},
	134: sysent{1, TF, "uselib", []int{ARG_STR, -1, -1, -1, -1, -1}},
	135: sysent{1, NF, "personality", []int{ARG_INT, -1, -1, -1, -1, -1}},
	136: sysent{2, 0, "ustat", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	137: sysent{2, TF, "statfs", []int{ARG_STR, ARG_PTR, -1, -1, -1, -1}},
	138: sysent{2, TD, "fstatfs", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	139: sysent{3, 0, "sysfs", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	140: sysent{2, 0, "getpriority", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	141: sysent{3, 0, "setpriority", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	142: sysent{2, TSC, "sched_setparam", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	143: sysent{2, TSC, "sched_getparam", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	144: sysent{3, TSC, "sched_setscheduler", []int{ARG_INT, ARG_INT, ARG_PTR, -1, -1, -1}},
	145: sysent{1, TSC, "sched_getscheduler", []int{ARG_INT, -1, -1, -1, -1, -1}},
	146: sysent{1, TSC, "sched_get_priority_max", []int{ARG_INT, -1, -1, -1, -1, -1}},
	147: sysent{1, TSC, "sched_get_priority_min", []int{ARG_INT, -1, -1, -1, -1, -1}},
	148: sysent{2, TSC, "sched_rr_get_interval", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	149: sysent{2, TM, "mlock", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	150: sysent{2, TM, "munlock", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	151: sysent{1, TM, "mlockall", []int{ARG_INT, -1, -1, -1, -1, -1}},
	152: sysent{0, TM, "munlockall", []int{-1, -1, -1, -1, -1, -1}},
	153: sysent{0, 0, "vhangup", []int{-1, -1, -1, -1, -1, -1}},
	154: sysent{3, 0, "modify_ldt", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	155: sysent{2, TF, "pivot_root", []int{ARG_STR, ARG_STR, -1, -1, -1, -1}},
	156: sysent{1, 0, "_sysctl", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	157: sysent{5, 0, "prctl", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1}},
	158: sysent{2, TP, "arch_prctl", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	159: sysent{1, 0, "adjtimex", []int{ARG_PTR, -1, -1, -1, -1, -1}},
	160: sysent{2, 0, "setrlimit", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	161: sysent{1, TF, "chroot", []int{ARG_STR, -1, -1, -1, -1, -1}},
	162: sysent{0, 0, "sync", []int{-1, -1, -1, -1, -1, -1}},
	163: sysent{1, TF, "acct", []int{ARG_STR, -1, -1, -1, -1, -1}},
	164: sysent{2, 0, "settimeofday", []int{ARG_PTR, ARG_PTR, -1, -1, -1, -1}},
	165: sysent{5, TF, "mount", []int{ARG_STR, ARG_STR, ARG_STR, ARG_INT, ARG_PTR, -1}},
	166: sysent{2, TF, "umount2", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	167: sysent{2, TF, "swapon", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	168: sysent{1, TF, "swapoff", []int{ARG_STR, -1, -1, -1, -1, -1}},
	169: sysent{4, 0, "reboot", []int{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, -1, -1}},
	170: sysent{2, 0, "sethostname", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	171: sysent{2, 0, "setdomainname", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	172: sysent{1, 0, "iopl", []int{ARG_INT, -1, -1, -1, -1, -1}},
	173: sysent{3, 0, "ioperm", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	174: sysent{2, 0, "create_module", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	175: sysent{3, 0, "init_module", []int{ARG_PTR, ARG_INT, ARG_STR, -1, -1, -1}},
	176: sysent{2, 0, "delete_module", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	177: sysent{1, 0, "get_kernel_syms", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	178: sysent{5, 0, "query_module", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	179: sysent{4, TF, "quotactl", []int{ARG_INT, ARG_STR, ARG_INT, ARG_PTR, -1, -1}},
	180: sysent{3, 0, "nfsservctl", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	181: sysent{5, TN, "getpmsg", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	182: sysent{5, TN, "putpmsg", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	183: sysent{5, 0, "afs_syscall", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	184: sysent{3, 0, "tuxcall", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	185: sysent{3, 0, "security", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	186: sysent{0, NF, "gettid", []int{-1, -1, -1, -1, -1, -1}},
	187: sysent{3, TD, "readahead", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	188: sysent{5, TF, "setxattr", []int{ARG_STR, ARG_STR, ARG_PTR, ARG_INT, ARG_INT, -1}},
	189: sysent{5, TF, "lsetxattr", []int{ARG_STR, ARG_STR, ARG_PTR, ARG_INT, ARG_INT, -1}},
	190: sysent{5, TD, "fsetxattr", []int{ARG_INT, ARG_STR, ARG_PTR, ARG_INT, ARG_INT, -1}},
	191: sysent{4, TF, "getxattr", []int{ARG_STR, ARG_STR, ARG_PTR, ARG_INT, -1, -1}},
	192: sysent{4, TF, "lgetxattr", []int{ARG_STR, ARG_STR, ARG_PTR, ARG_INT, -1, -1}},
	193: sysent{4, TD, "fgetxattr", []int{ARG_INT, ARG_STR, ARG_PTR, ARG_INT, -1, -1}},
	194: sysent{3, TF, "listxattr", []int{ARG_STR, ARG_STR, ARG_INT, -1, -1, -1}},
	195: sysent{3, TF, "llistxattr", []int{ARG_STR, ARG_STR, ARG_INT, -1, -1, -1}},
	196: sysent{3, TD, "flistxattr", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	197: sysent{2, TF, "removexattr", []int{ARG_STR, ARG_STR, -1, -1, -1, -1}},
	198: sysent{2, TF, "lremovexattr", []int{ARG_STR, ARG_STR, -1, -1, -1, -1}},
	199: sysent{2, TD, "fremovexattr", []int{ARG_INT, ARG_STR, -1, -1, -1, -1}},
	200: sysent{2, TS, "tkill", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	201: sysent{1, 0, "time", []int{ARG_PTR, -1, -1, -1, -1, -1}},
	202: sysent{6, 0, "futex", []int{ARG_PTR, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}},
	203: sysent{3, TSC, "sched_setaffinity", []int{ARG_INT, ARG_INT, ARG_PTR, -1, -1, -1}},
	204: sysent{3, TSC, "sched_getaffinity", []int{ARG_INT, ARG_INT, ARG_PTR, -1, -1, -1}},
	205: sysent{1, 0, "set_thread_area", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	206: sysent{2, TM, "io_setup", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	207: sysent{1, TM, "io_destroy", []int{ARG_INT, -1, -1, -1, -1, -1}},
	208: sysent{5, 0, "io_getevents", []int{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, -1}},
	209: sysent{3, 0, "io_submit", []int{ARG_INT, ARG_INT, ARG_PTR, -1, -1, -1}},
	210: sysent{3, 0, "io_cancel", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	211: sysent{1, 0, "get_thread_area", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	212: sysent{3, 0, "lookup_dcookie", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	213: sysent{1, TD, "epoll_create", []int{ARG_INT, -1, -1, -1, -1, -1}},
	214: sysent{4, 0, "epoll_ctl_old", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	215: sysent{4, 0, "epoll_wait_old", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	216: sysent{5, TM | SI, "remap_file_pages", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1}},
	217: sysent{3, TD, "getdents64", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	218: sysent{1, 0, "set_tid_address", []int{ARG_PTR, -1, -1, -1, -1, -1}},
	219: sysent{0, 0, "restart_syscall", []int{-1, -1, -1, -1, -1, -1}},
	220: sysent{4, TI, "semtimedop", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, -1, -1}},
	221: sysent{4, TD, "fadvise64", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1, -1}},
	222: sysent{3, 0, "timer_create", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	223: sysent{4, 0, "timer_settime", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, -1, -1}},
	224: sysent{2, 0, "timer_gettime", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	225: sysent{1, 0, "timer_getoverrun", []int{ARG_INT, -1, -1, -1, -1, -1}},
	226: sysent{1, 0, "timer_delete", []int{ARG_INT, -1, -1, -1, -1, -1}},
	227: sysent{2, 0, "clock_settime", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	228: sysent{2, 0, "clock_gettime", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	229: sysent{2, 0, "clock_getres", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	230: sysent{4, 0, "clock_nanosleep", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, -1, -1}},
	231: sysent{1, TP | SE, "exit_group", []int{ARG_INT, -1, -1, -1, -1, -1}},
	232: sysent{4, TD, "epoll_wait", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, -1, -1}},
	233: sysent{4, TD, "epoll_ctl", []int{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, -1, -1}},
	234: sysent{3, TS, "tgkill", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	235: sysent{2, TF, "utimes", []int{ARG_STR, ARG_PTR, -1, -1, -1, -1}},
	236: sysent{5, 0, "vserver", []int{ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	237: sysent{6, TM, "mbind", []int{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}},
	238: sysent{3, TM, "set_mempolicy", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	239: sysent{5, TM, "get_mempolicy", []int{ARG_PTR, ARG_PTR, ARG_INT, ARG_INT, ARG_INT, -1}},
	240: sysent{4, 0, "mq_open", []int{ARG_STR, ARG_INT, ARG_INT, ARG_PTR, -1, -1}},
	241: sysent{1, 0, "mq_unlink", []int{ARG_STR, -1, -1, -1, -1, -1}},
	242: sysent{5, 0, "mq_timedsend", []int{ARG_INT, ARG_STR, ARG_INT, ARG_INT, ARG_PTR, -1}},
	243: sysent{5, 0, "mq_timedreceive", []int{ARG_INT, ARG_STR, ARG_INT, ARG_PTR, ARG_PTR, -1}},
	244: sysent{2, 0, "mq_notify", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	245: sysent{3, 0, "mq_getsetattr", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	246: sysent{4, 0, "kexec_load", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_INT, -1, -1}},
	247: sysent{5, TP, "waitid", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, -1}},
	248: sysent{5, 0, "add_key", []int{ARG_STR, ARG_STR, ARG_PTR, ARG_INT, ARG_INT, -1}},
	249: sysent{4, 0, "request_key", []int{ARG_STR, ARG_STR, ARG_STR, ARG_INT, -1, -1}},
	250: sysent{5, 0, "keyctl", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1}},
	251: sysent{3, 0, "ioprio_set", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	252: sysent{2, 0, "ioprio_get", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	253: sysent{0, TD, "inotify_init", []int{-1, -1, -1, -1, -1, -1}},
	254: sysent{3, TD, "inotify_add_watch", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	255: sysent{2, TD, "inotify_rm_watch", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	256: sysent{4, TM, "migrate_pages", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, -1, -1}},
	257: sysent{4, TD | TF, "openat", []int{ARG_INT, ARG_STR, ARG_INT, ARG_INT, -1, -1}},
	258: sysent{3, TD | TF, "mkdirat", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	259: sysent{4, TD | TF, "mknodat", []int{ARG_INT, ARG_STR, ARG_INT, ARG_INT, -1, -1}},
	260: sysent{5, TD | TF, "fchownat", []int{ARG_INT, ARG_STR, ARG_INT, ARG_INT, ARG_INT, -1}},
	261: sysent{3, TD | TF, "futimesat", []int{ARG_INT, ARG_STR, ARG_PTR, -1, -1, -1}},
	262: sysent{4, TD | TF, "newfstatat", []int{ARG_INT, ARG_STR, ARG_PTR, ARG_INT, -1, -1}},
	263: sysent{3, TD | TF, "unlinkat", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	264: sysent{4, TD | TF, "renameat", []int{ARG_INT, ARG_STR, ARG_INT, ARG_STR, -1, -1}},
	265: sysent{5, TD | TF, "linkat", []int{ARG_INT, ARG_STR, ARG_INT, ARG_STR, ARG_INT, -1}},
	266: sysent{3, TD | TF, "symlinkat", []int{ARG_STR, ARG_INT, ARG_STR, -1, -1, -1}},
	267: sysent{4, TD | TF, "readlinkat", []int{ARG_INT, ARG_STR, ARG_STR, ARG_INT, -1, -1}},
	268: sysent{3, TD | TF, "fchmodat", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	269: sysent{3, TD | TF, "faccessat", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	270: sysent{6, TD, "pselect6", []int{ARG_INT, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}},
	271: sysent{5, TD, "ppoll", []int{ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR, ARG_INT, -1}},
	272: sysent{1, TP, "unshare", []int{ARG_INT, -1, -1, -1, -1, -1}},
	273: sysent{2, 0, "set_robust_list", []int{ARG_PTR, ARG_INT, -1, -1, -1, -1}},
	274: sysent{3, 0, "get_robust_list", []int{ARG_INT, ARG_PTR, ARG_PTR, -1, -1, -1}},
	275: sysent{6, TD, "splice", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}},
	276: sysent{4, TD, "tee", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1, -1}},
	277: sysent{4, TD, "sync_file_range", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1, -1}},
	278: sysent{4, TD, "vmsplice", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, -1, -1}},
	279: sysent{6, TM, "move_pages", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, ARG_PTR, ARG_INT}},
	280: sysent{4, TD | TF, "utimensat", []int{ARG_INT, ARG_STR, ARG_PTR, ARG_INT, -1, -1}},
	281: sysent{6, TD, "epoll_pwait", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, ARG_PTR, ARG_INT}},
	282: sysent{3, TD | TS, "signalfd", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	283: sysent{2, TD, "timerfd_create", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	284: sysent{1, TD, "eventfd", []int{ARG_INT, -1, -1, -1, -1, -1}},
	285: sysent{4, TD, "fallocate", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1, -1}},
	286: sysent{4, TD, "timerfd_settime", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, -1, -1}},
	287: sysent{2, TD, "timerfd_gettime", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	288: sysent{4, TN, "accept4", []int{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT, -1, -1}},
	289: sysent{4, TD | TS, "signalfd4", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, -1, -1}},
	290: sysent{2, TD, "eventfd2", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	291: sysent{1, TD, "epoll_create1", []int{ARG_INT, -1, -1, -1, -1, -1}},
	292: sysent{3, TD, "dup3", []int{ARG_INT, ARG_INT, ARG_INT, -1, -1, -1}},
	293: sysent{2, TD, "pipe2", []int{ARG_PTR, ARG_INT, -1, -1, -1, -1}},
	294: sysent{1, TD, "inotify_init1", []int{ARG_INT, -1, -1, -1, -1, -1}},
	295: sysent{4, TD, "preadv", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, ARG_INT, -1}},
	296: sysent{4, TD, "pwritev", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, ARG_INT, -1}},
	297: sysent{4, TP | TS, "rt_tgsigqueueinfo", []int{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, -1, -1}},
	298: sysent{5, TD, "perf_event_open", []int{ARG_PTR, ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1}},
	299: sysent{5, TN, "recvmmsg", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, ARG_PTR, -1}},
	300: sysent{2, TD, "fanotify_init", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	301: sysent{5, TD | TF, "fanotify_mark", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_STR, -1}},
	302: sysent{4, 0, "prlimit64", []int{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR, -1, -1}},
	303: sysent{5, TD | TF, "name_to_handle_at", []int{ARG_INT, ARG_STR, ARG_PTR, ARG_PTR, ARG_INT, -1}},
	304: sysent{3, TD, "open_by_handle_at", []int{ARG_INT, ARG_PTR, ARG_INT, -1, -1, -1}},
	305: sysent{2, 0, "clock_adjtime", []int{ARG_INT, ARG_PTR, -1, -1, -1, -1}},
	306: sysent{1, TD, "syncfs", []int{ARG_INT, -1, -1, -1, -1, -1}},
	307: sysent{4, TN, "sendmmsg", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_INT, -1, -1}},
	308: sysent{2, TD, "setns", []int{ARG_INT, ARG_INT, -1, -1, -1, -1}},
	309: sysent{3, 0, "getcpu", []int{ARG_PTR, ARG_PTR, ARG_PTR, -1, -1, -1}},
	310: sysent{6, 0, "process_vm_readv", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}},
	311: sysent{6, 0, "process_vm_writev", []int{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}},
	312: sysent{5, 0, "kcmp", []int{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, -1}},
	313: sysent{3, TD, "finit_module", []int{ARG_INT, ARG_STR, ARG_INT, -1, -1, -1}},
	314: sysent{3, TSC, "sched_setattr", []int{}},
	315: sysent{4, TSC, "sched_getattr", []int{}},
	316: sysent{5, TD | TF, "renameat2", []int{ARG_INT, ARG_STR, ARG_INT, ARG_STR, ARG_INT, -1}},
	317: sysent{3, 0, "seccomp", []int{}},
	318: sysent{3, 0, "getrandom", []int{}},
	319: sysent{2, TD, "memfd_create", []int{ARG_STR, ARG_INT, -1, -1, -1, -1}},
	320: sysent{5, TD, "kexec_file_load", []int{ARG_INT, ARG_INT, ARG_INT, ARG_STR, ARG_INT, -1}},
	321: sysent{3, TD, "bpf", []int{}},
	322: sysent{5, TD | TF | TP | SE | SI, "execveat", []int{}},
	323: sysent{1, TD, "userfaultfd", []int{ARG_INT, -1, -1, -1, -1, -1}},
	324: sysent{2, 0, "membarrier", []int{}},
	325: sysent{3, TM, "mlock2", []int{}},
	326: sysent{6, TD, "copy_file_range", []int{}},
	327: sysent{6, TD, "preadv2", []int{}},
	328: sysent{6, TD, "pwritev2", []int{}},
	329: sysent{4, TM | SI, "pkey_mprotect", []int{}},
	330: sysent{2, 0, "pkey_alloc", []int{}},
	331: sysent{1, 0, "pkey_free", []int{}},
	332: sysent{5, TD | TF, "statx", []int{}},
}
