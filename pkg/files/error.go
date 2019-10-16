// Copyright 2019 Vikunja and contriubtors. All rights reserved.
//
// This file is part of Vikunja.
//
// Vikunja is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Vikunja is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Vikunja.  If not, see <https://www.gnu.org/licenses/>.

package files

import "fmt"

// ErrFileDoesNotExist defines an error where a file does not exist in the db
type ErrFileDoesNotExist struct {
	FileID int64
}

// Error is the error implementation of ErrFileDoesNotExist
func (err ErrFileDoesNotExist) Error() string {
	return fmt.Sprintf("file %d does not exist", err.FileID)
}

//IsErrFileDoesNotExist checks if an error is ErrFileDoesNotExist
func IsErrFileDoesNotExist(err error) bool {
	_, ok := err.(ErrFileDoesNotExist)
	return ok
}

// ErrFileIsTooLarge defines an error where a file is larger than the configured limit
type ErrFileIsTooLarge struct {
	Size int64
}

// Error is the error implementation of ErrFileIsTooLarge
func (err ErrFileIsTooLarge) Error() string {
	return fmt.Sprintf("file is too large [Size: %d]", err.Size)
}

//IsErrFileIsTooLarge checks if an error is ErrFileIsTooLarge
func IsErrFileIsTooLarge(err error) bool {
	_, ok := err.(ErrFileIsTooLarge)
	return ok
}
