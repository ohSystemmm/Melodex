package Music

//func GetSongLength(song []byte) (int, error) {
//	reader := bytes.NewReader(song)
//
//	streamer, format, err := mp3.Decode(reader)
//	if err != nil {
//		return 0, err
//	}
//	defer streamer.Close()
//
//	duration := time.Duration(streamer.Len()) * time.Second / time.Duration(format.SampleRate)
//
//	return int(duration.Seconds()), nil
//}
