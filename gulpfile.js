var gulp = require('gulp');
var concat = require('gulp-concat')
var uglify = require('gulp-uglify')
var minifyCSS = require('gulp-minify-css')
var strip = require('gulp-strip-comments')

var dependencies_path = './node_modules/';
var public_path = 'web/public';

gulp.task('minimize_fastclick', function() {

    return gulp.src(dependencies_path + 'fastclick/lib/fastclick.js')
        .pipe(uglify())
        .pipe(gulp.dest(dependencies_path + 'fastclick/lib'));

});

gulp.task('compile_js', ['minimize_fastclick'], function() {


    return gulp.src(
            [
                dependencies_path + 'jquery/dist/jquery.min.js',
                dependencies_path + 'bootstrap/dist/js/bootstrap.min.js',
                dependencies_path + 'fastclick/lib/fastclick.js',
                dependencies_path + 'admin-lte/dist/js/adminlte.min.js',
                dependencies_path + 'angular/angular.min.js'
            ]
        )
        .pipe(strip.text())
        .pipe(uglify())
        .pipe(concat('js/app.min.js'))
        .pipe(gulp.dest(public_path));

});

gulp.task('minimize_datatables_bootstrap', function() {

    return gulp.src(dependencies_path + 'datatables.net-bs/css/dataTables.bootstrap.css')
        .pipe(minifyCSS())
        .pipe(gulp.dest(dependencies_path + 'datatables.net-bs/css'));

});

gulp.task('compile_css', ['minimize_datatables_bootstrap'], function() {


    return gulp.src(
            [
                dependencies_path + 'bootstrap/dist/css/bootstrap.min.css',
                dependencies_path + 'font-awesome/css/font-awesome.min.css',
                dependencies_path + 'ionicons/css/ionicons.min.css',
                dependencies_path + 'datatables.net-bs/css/dataTables.bootstrap.css',
                dependencies_path + 'admin-lte/dist/css/AdminLTE.min.css',
                dependencies_path + 'admin-lte/dist/css/skins/_all-skins.min.css'
            ]
        )
        .pipe(strip.text())
        .pipe(concat('css/app.min.css'))
        .pipe(gulp.dest(public_path));

});

gulp.task('copile_assets', ['compile_js', 'compile_css']);

gulp.task('run', ['copile_assets']);