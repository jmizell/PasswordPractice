# Password Practice Tool

This tool helps users securely practice typing infrequently used passwords. It
allows someone to remember important passwords that they may not use very
often, in an easy-to-read and secure manner.

## Installation

1. Install [Go](https://golang.org/doc/install) if you haven't already.
2. Install the required dependencies:

```
go get -u github.com/howeyc/gopass
go get -u golang.org/x/crypto/bcrypt
```

3. Compile the program:

```
go build -o password-practice
```

## Usage

### Adding a password

To add a new password to the config file:

```
./password-practice -add
```

You will be prompted to enter the account name and the password you'd like to
store. The password will be securely hashed before storing it in the config
file.

### Practicing passwords

To practice typing passwords:

```
./password-practice
```

The program will loop through each account in the config file and prompt you
to enter the correct password. If the entered password is correct, the tool
will proceed to the next account. If the entered password is incorrect, it
will prompt you again to re-enter the password.

## Configuration

By default, the tool uses a config file named `config.json`. If you want to
use a different file, you can provide the `-config` flag with the desired file
path:

```
./password-practice -config myconfig.json
```

This flag works with both adding and practicing passwords.

## Security

The tool uses the bcrypt password hashing algorithm to securely store passwords
in the config file. Passwords are never stored in plaintext.

When entering passwords, the tool uses the `gopass` library to read input
without echoing it on the screen.

## License

This tool is provided under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html).

## Disclaimer

Please note that the security of your passwords depends on the security of the
machine where the config file is stored. It is your responsibility to ensure
the config file is not accessible by unauthorized parties.