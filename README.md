# File Encryption & Decryption Tool

This GoLang project uses AED 256 encryption to secure files. Built with Cobra, it provides simple CLI commands for encryption and decryption.

## Usage

- **Encrypt:** `encrypt -f /path/to/file -p "password"`
- **Decrypt:** `decrypt -f /path/to/encrypted_file.aed -p "password"`

Example:

- Encrypt: `encrypt -f /george.txt -p "welcome"` -> Output: `george.aed`
- Decrypt: `decrypt -f /george.aed -p "welcome"` -> Output: `george.txt`

## License

MIT License. Contributions welcome.

