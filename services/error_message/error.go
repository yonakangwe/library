/**
 * @author Yohana Kangwe
 * @email yonakangwe@gmail.com
 * @create date 2024-03-20 00:53:05
 * @modify date 2024-03-20 00:53:05
 * @desc [description]
 */

package error_message

import "errors"

var (
	ErrNoResultSet    = errors.New("no rows in result set")
	ErrDuplicateEntry = errors.New("record already exists")
)
