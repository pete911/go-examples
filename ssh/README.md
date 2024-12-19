Validate SSH public key

```go
import "golang.org/x/crypto/ssh"

func ValidatePublicKey(publicKey string) error {
    if _, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey)); err != nil {
        return err
    }
    return nil
}

```
