package main

import "time"

type User struct {
    ID                 int        `json:"id"`
    Email              string     `json:"email"`
    Username           string     `json:"username"`
    FirstName          string     `json:"first_name"`
    LastName           string     `json:"last_name"`
    PasswordHash       string     `json:"password_hash"`
    IsActive           bool       `json:"is_active"`
    IsVerified         bool       `json:"is_verified"`
    VerificationToken  string     `json:"verification_token"`
    ResetToken         string     `json:"reset_token"`
    ResetTokenExpires  *time.Time `json:"reset_token_expires"`
    FailedLoginAttempts int       `json:"failed_login_attempts"`
    LockUntil          *time.Time `json:"lock_until"`
    LastLogin          *time.Time `json:"last_login"`
    CreatedAt          time.Time  `json:"created_at"`
    UpdatedAt          time.Time  `json:"updated_at"`
    DeletedAt          *time.Time `json:"deleted_at"`
}

