package day22

const (
	steps       = 2000
	pruneFactor = 0xffffff
)

func calculateNextSecret(secret int) int {
	// multiply by 64 and prune last 24 bits (modulo 16777214)
	secret ^= (secret << 6) & pruneFactor
	// divide by 32 and prune last 24 bits (modulo 16777214)
	secret ^= (secret >> 5) & pruneFactor
	// multiply by 2048 and prune last 24 bits (modulo 16777214)
	secret ^= (secret << 11) & pruneFactor
	return secret
}
