// pls (c) by <Ryan Rasti - rrrasti@yahoo.com>

// pls is licensed under a
// Creative Commons Attribution-NonCommercial-NoDerivatives 4.0 International License.

// You should have received a copy of the license along with this
// work. If not, see <http://creativecommons.org/licenses/by-nc-nd/4.0/>.

package main

import "github.com/diamonddelt/pls/cmd"

var (
	// VERSION should be set by CI for versioning the artifact
	VERSION = "0.0.3"
)

func main() {
	cmd.Execute(VERSION)
}
