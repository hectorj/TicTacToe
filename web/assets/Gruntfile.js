module.exports = function(grunt) {

    // Project configuration.
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        less: {
            development: {
                options: {
                    compress: false,
                },
                files: {
                    "dist/main.css": "main.less"
                }
            },
            production: {
                options: {
                    compress: true,
                },
                files: {
                    "dist/main.css": "main.less"
                }
            }
        }
    });

    grunt.loadNpmTasks('grunt-contrib-less');

    // Default task(s).
    grunt.registerTask('default', ['less:production']);

};