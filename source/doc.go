// Package source holds all the assets for Statping. This includes
// CSS, JS, SCSS, HTML and other website related content.
// This package uses Rice to compile all assets into a single 'rice-box.go' file.
//
// Required Dependencies
//
// - rice -> https://github.com/GeertJohan/go.rice
// - sass -> https://sass-lang.com/install
//
// Compile Assets
//
// To compile all the HTML, JS, SCSS, CSS and image assets you'll need to have rice and sass installed on your local system.
//		sass source/scss/base.scss source/css/base.css
//		cd source && rice embed-go
//
// More info on: https://github.com/hunterlong/statping
package source
