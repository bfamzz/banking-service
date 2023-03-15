package db

import "context"

type VerifyEmailTxParams struct {
	VerifyEmailParams
	AfterVerify func(verifyEmail VerifyEmail) error
}

type VerifyEmailTxResult struct {
	VerifyEmail VerifyEmail
}

func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.VerifyEmail(ctx, arg.VerifyEmailParams)
		if err != nil {
			return err
		}

		return arg.AfterVerify(result.VerifyEmail)
	})

	return result, err
}
