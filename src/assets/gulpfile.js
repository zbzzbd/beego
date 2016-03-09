var gulp = require('gulp'),
    minifycss = require('gulp-minify-css'),
    uglify = require('gulp-uglify'),
    rename = require('gulp-rename'),
    clean = require('gulp-clean'),
    concat = require('gulp-concat'),
    livereload = require('gulp-livereload');


var path = function(){
    var base =  "../public/";
    return {
        js: base + "js/",
        css : base + "css/",
        img : base + "img/"
    }
}();

console.log(path)

gulp.task('clean', function() {
    return gulp.src([path.css + '*.css', path.js + '*.js', path.img + "*"], {read: false})
        .pipe(clean({force:true}));
});

gulp.task('css', function() {
    return gulp.src([
        'css/libs/**.css',
        '!css/libs/semantic.min.css',
        'css/**.css',
        'css/main.css'
    ])
        .pipe(concat('main.css'))
        .pipe(gulp.dest(path.css))
        .pipe(rename({ suffix: '.min' }))
        .pipe(minifycss())
        .pipe(gulp.dest(path.css))
        .pipe(livereload());
});

gulp.task('libjs', function() {
    return gulp.src([
        'node_modules/jquery/dist/jquery.js',
        'node_modules/jquery.cookie/jquery.cookie.js',
        'node_modules/imagesloaded/imagesloaded.pkgd.js'
        ])
        .pipe(concat('lib.js'))
        .pipe(gulp.dest(path.js))
        .pipe(rename({suffix: '.min'}))
        .pipe(uglify())
        .pipe(gulp.dest(path.js))
        .pipe(livereload());
});

gulp.task("appjs", function(){
return gulp.src([
        'js/libs/*.js',
        'js/**/*.js',
        '!js/dev/**/*.js'
    ])
    .pipe(concat('app.js'))  
    .pipe(gulp.dest(path.js))
    .pipe(rename({suffix: '.min'}))
    .pipe(uglify())
    .pipe(gulp.dest(path.js))
    .pipe(livereload());
});

gulp.task('jsDev', function() {
    return gulp.src([
        'node_modules/jquery-mockjax/dist/jquery.mockjax.js',
        'js/dev/**/*.js'
    ])
        .pipe(concat('dev.js'))
        .pipe(gulp.dest(path.js))
        .pipe(rename({suffix: '.min'}))
        .pipe(uglify())
        .pipe(gulp.dest(path.js))
        .pipe(livereload());
});

gulp.task('cpSemanticCss', function() {
    return gulp.src([
        'css/libs/semantic.min.css',
        'node_modules/semantic-ui/dist/**',
        '!node_modules/semantic-ui/dist/*.css',
        '!node_modules/semantic-ui/dist/components/**',
        '!node_modules/semantic-ui/dist/*.js'
    ])
        .pipe(gulp.dest(path.css));
});

gulp.task('cpSemanticJs', function(){
    return gulp.src([
        'node_modules/semantic-ui/dist/*.js',
    ])
    .pipe(gulp.dest(path.js))
});

gulp.task("cpimg", function(){
    return gulp.src([
        "img/*"
    ])
    .pipe(gulp.dest(path.img))
});


gulp.task('default', ['clean'], function() {
    gulp.start('css', 'libjs', 'appjs','jsDev', 'cpSemanticCss','cpSemanticJs', 'cpimg');
});


gulp.task('watch', function() {
    gulp.watch("img/**", ['cpimg'])
    gulp.watch('css/**/*.css', ['css']);

    gulp.watch([
        'js/**/*.js',
        'js/*.js',
        '!js/libs/html5.js',
        '!js/dev/**/*.js'
    ], ['appjs']);

    gulp.watch([
        'js/dev/**/*.js'
    ], ['jsDev']);

    livereload.listen();
    gulp.watch(['dist/**']).on('change', function(file) {
        livereload.changed(file.path);
    });
});
