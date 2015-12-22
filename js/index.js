// Imports and uses the 'js.js' generated through gopherjs
'use strict';
console.log('Starting');

console.log('Loading js.js');

require('./js.js');

var mzika = global.mzika;
var vid = "uscmv1500002";
console.log('Looking up a video with id:', vid);
var video_url = mzika.getVideoUrl(vid);
console.log('Got Video URL: ', video_url);

console.log('Fetching Full Video Struct Object');
var video_struct = mzika.getVideoFromId(vid);
console.log('Got Video Struct: ', video_struct);

console.log('Done!');
