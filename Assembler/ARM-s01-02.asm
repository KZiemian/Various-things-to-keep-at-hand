# ARM Assembly Tutorial s01-02

# File descryptors
# STDIN:  0
# STDOUT: 1
# STDERR: 2

# Write a string to stdout file descryptor.

.global _start
.section .text



_start:
	mov r7, #0x4
	mov r0, #1
	ldr r1, =message
	mov r2, #14

	swi 0

	mov r7, #0x1
	mov r0, #65

	swi 0



.section .data
message:
	.ascii "Hello, World!\n"
