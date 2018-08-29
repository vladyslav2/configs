package main

import (
	"fmt"
	"time"

	"github.com/bitmovin/bitmovin-go/bitmovin"
	"github.com/bitmovin/bitmovin-go/bitmovintypes"
	"github.com/bitmovin/bitmovin-go/models"
	"github.com/bitmovin/bitmovin-go/services"
)

func main() {

	inputPath := "/demos/sample-videos/small.mp4"
	outputPath := "/Exampledirectory"

	// create a bitmovin api client
	bitmovin := bitmovin.NewBitmovin("5ab9487b-bb1b-44f0-9772-7ef4b3b62288", "https://api.bitmovin.com/v1/", 5)

	// create the https input
	httpIS := services.NewHTTPInputService(bitmovin)
	httpsInput := &models.HTTPInput{
		Host: stringToPtr("http://techslides.com"),
	}
	httpsResp, err := httpIS.Create(httpsInput)
	errorHandler(httpsResp.Status, err)

	// create the s3 output
	/*
		s3OS := services.NewS3OutputService(bitmovin)
		s3Output := &models.S3Output{
			AccessKey:   stringToPtr("your-s3-access-key"),
			SecretKey:   stringToPtr("your-s3-secret-key"),
			BucketName:  stringToPtr("your-s3-bucket-name"),
			CloudRegion: bitmovintypes.AWSCloudRegionEUWest1,
		}
		s3OutputResp, err := s3OS.Create(s3Output)
		errorHandler(s3OutputResp.Status, err)
	*/
	outputSer := services.NewGCSOutputService(bitmovin)
	s3OutputResp, err := outputSer.Retrieve("3785f414-1a72-4c45-9329-5e3e592d8dd9")
	errorHandler(s3OutputResp.Status, err)

	// create the video and audio codec configurations
	h264S := services.NewH264CodecConfigurationService(bitmovin)
	video1080pConfig := &models.H264CodecConfiguration{
		Name:    stringToPtr("my-h264-4-8mbit-1080p-cc"),
		Bitrate: intToPtr(4800000),
		Height:  intToPtr(1080),
		Profile: bitmovintypes.H264ProfileHigh,
	}

	video1080pResp, err := h264S.Create(video1080pConfig)
	errorHandler(video1080pResp.Status, err)

	aacS := services.NewAACCodecConfigurationService(bitmovin)
	aacConfig := &models.AACCodecConfiguration{
		Name:    stringToPtr("my-aac-128kbit-cc"),
		Bitrate: intToPtr(128000),
	}
	aacResp, err := aacS.Create(aacConfig)
	errorHandler(aacResp.Status, err)

	// create the encoding resource
	encodingS := services.NewEncodingService(bitmovin)
	encoding := &models.Encoding{
		Name:        stringToPtr("my-awesome-second-encoding"),
		CloudRegion: bitmovintypes.CloudRegionGoogleEuropeWest1,
	}
	encodingResp, err := encodingS.Create(encoding)
	errorHandler(encodingResp.Status, err)

	// add the video and audio streams to the encoding
	videoInputStream := models.InputStream{
		InputID:       httpsResp.Data.Result.ID,
		InputPath:     stringToPtr(inputPath),
		SelectionMode: bitmovintypes.SelectionModeAuto,
	}

	audioInputStream := models.InputStream{
		InputID:       httpsResp.Data.Result.ID,
		InputPath:     stringToPtr(inputPath),
		SelectionMode: bitmovintypes.SelectionModeAuto,
	}

	vis := []models.InputStream{videoInputStream}
	videoStream1080p := &models.Stream{
		CodecConfigurationID: video1080pResp.Data.Result.ID,
		InputStreams:         vis,
	}
	videoStream1080pResp, err := encodingS.AddStream(*encodingResp.Data.Result.ID, videoStream1080p)
	errorHandler(videoStream1080pResp.Status, err)

	ais := []models.InputStream{audioInputStream}
	audioStream := &models.Stream{
		CodecConfigurationID: aacResp.Data.Result.ID,
		InputStreams:         ais,
	}
	aacStreamResp, err := encodingS.AddStream(*encodingResp.Data.Result.ID, audioStream)
	errorHandler(aacStreamResp.Status, err)

	// add the video muxings to the encoding
	aclEntry := models.ACLItem{
		Permission: bitmovintypes.ACLPermissionPublicRead,
	}

	acl := []models.ACLItem{aclEntry}

	videoMuxingStream1080p := models.StreamItem{
		StreamID: videoStream1080pResp.Data.Result.ID,
	}

	videoMuxingFmp41080pOutput := models.Output{
		OutputID:   s3OutputResp.Data.Result.ID,
		OutputPath: stringToPtr(outputPath + "/video/1080p/fmp4"),
		ACL:        acl,
	}

	videoMuxingTs1080pOutput := models.Output{
		OutputID:   s3OutputResp.Data.Result.ID,
		OutputPath: stringToPtr(outputPath + "/video/1080p/ts"),
		ACL:        acl,
	}

	videoFMP4Muxing1080p := &models.FMP4Muxing{
		SegmentLength:   floatToPtr(4.0),
		SegmentNaming:   stringToPtr("seg_%number%.m4s"),
		InitSegmentName: stringToPtr("init.mp4"),
		Streams:         []models.StreamItem{videoMuxingStream1080p},
		Outputs:         []models.Output{videoMuxingFmp41080pOutput},
	}

	videoFMP4Muxing1080pResp, err := encodingS.AddFMP4Muxing(*encodingResp.Data.Result.ID, videoFMP4Muxing1080p)
	errorHandler(videoFMP4Muxing1080pResp.Status, err)

	videoTSMuxing1080p := &models.TSMuxing{
		SegmentLength: floatToPtr(4.0),
		SegmentNaming: stringToPtr("seg_%number%.ts"),
		Streams:       []models.StreamItem{videoMuxingStream1080p},
		Outputs:       []models.Output{videoMuxingTs1080pOutput},
	}

	videoTSMuxing1080pResp, err := encodingS.AddTSMuxing(*encodingResp.Data.Result.ID, videoTSMuxing1080p)
	errorHandler(videoTSMuxing1080pResp.Status, err)

	// add the audio muxings to the encoding

	audioMuxingStream := models.StreamItem{
		StreamID: aacStreamResp.Data.Result.ID,
	}

	audioMuxingFmp4Output := models.Output{
		OutputID:   s3OutputResp.Data.Result.ID,
		OutputPath: stringToPtr(outputPath + "/audio/fmp4"),
		ACL:        acl,
	}

	audioMuxingTsOutput := models.Output{
		OutputID:   s3OutputResp.Data.Result.ID,
		OutputPath: stringToPtr(outputPath + "/audio/ts"),
		ACL:        acl,
	}

	audioFMP4Muxing := &models.FMP4Muxing{
		SegmentLength:   floatToPtr(4.0),
		SegmentNaming:   stringToPtr("seg_%number%.m4s"),
		InitSegmentName: stringToPtr("init.mp4"),
		Streams:         []models.StreamItem{audioMuxingStream},
		Outputs:         []models.Output{audioMuxingFmp4Output},
	}
	audioFMP4MuxingResp, err := encodingS.AddFMP4Muxing(*encodingResp.Data.Result.ID, audioFMP4Muxing)
	errorHandler(audioFMP4MuxingResp.Status, err)

	audioTSMuxing := &models.TSMuxing{
		SegmentLength: floatToPtr(4.0),
		SegmentNaming: stringToPtr("seg_%number%.ts"),
		Streams:       []models.StreamItem{audioMuxingStream},
		Outputs:       []models.Output{audioMuxingTsOutput},
	}
	audioTSMuxingResp, err := encodingS.AddTSMuxing(*encodingResp.Data.Result.ID, audioTSMuxing)
	errorHandler(audioTSMuxingResp.Status, err)

	// start the encoding

	startResp, err := encodingS.Start(*encodingResp.Data.Result.ID)
	errorHandler(startResp.Status, err)

	waitForEncoding(encodingS, encodingResp)

	// create the dash manifest

	manifestOutput := models.Output{
		OutputID:   s3OutputResp.Data.Result.ID,
		OutputPath: stringToPtr(outputPath),
		ACL:        acl,
	}
	dashManifest := &models.DashManifest{
		Name:         stringToPtr("my-awesome-dash-manifest"),
		ManifestName: stringToPtr("manifest.mpd"),
		Outputs:      []models.Output{manifestOutput},
	}

	dashService := services.NewDashManifestService(bitmovin)
	dashManifestResp, err := dashService.Create(dashManifest)
	errorHandler(dashManifestResp.Status, err)

	period := &models.Period{}
	periodResp, err := dashService.AddPeriod(*dashManifestResp.Data.Result.ID, period)
	errorHandler(periodResp.Status, err)

	// add the video and audio adaptation sets to the period

	vas := &models.VideoAdaptationSet{}
	vasResp, err := dashService.AddVideoAdaptationSet(*dashManifestResp.Data.Result.ID, *periodResp.Data.Result.ID, vas)
	errorHandler(vasResp.Status, err)

	aas := &models.AudioAdaptationSet{
		Language: stringToPtr("en"),
	}
	aasResp, err := dashService.AddAudioAdaptationSet(*dashManifestResp.Data.Result.ID, *periodResp.Data.Result.ID, aas)
	errorHandler(aasResp.Status, err)

	// add the fmp4 representations to the adaptation sets

	fmp4Rep1080 := &models.FMP4Representation{
		Type:        bitmovintypes.FMP4RepresentationTypeTemplate,
		MuxingID:    videoFMP4Muxing1080pResp.Data.Result.ID,
		EncodingID:  encodingResp.Data.Result.ID,
		SegmentPath: stringToPtr(outputPath + "/video/1080p/fmp4"),
	}
	fmp4Rep1080Resp, err := dashService.AddFMP4Representation(*dashManifestResp.Data.Result.ID, *periodResp.Data.Result.ID, *vasResp.Data.Result.ID, fmp4Rep1080)
	errorHandler(fmp4Rep1080Resp.Status, err)

	fmp4RepAudio := &models.FMP4Representation{
		Type:        bitmovintypes.FMP4RepresentationTypeTemplate,
		MuxingID:    audioFMP4MuxingResp.Data.Result.ID,
		EncodingID:  encodingResp.Data.Result.ID,
		SegmentPath: stringToPtr(outputPath + "/audio/fmp4"),
	}
	fmp4RepAudioResp, err := dashService.AddFMP4Representation(*dashManifestResp.Data.Result.ID, *periodResp.Data.Result.ID, *aasResp.Data.Result.ID, fmp4RepAudio)
	errorHandler(fmp4RepAudioResp.Status, err)

	// start the dash manifest generation

	startResp, err = dashService.Start(*dashManifestResp.Data.Result.ID)
	errorHandler(startResp.Status, err)

	waitForDashManifest(dashService, dashManifestResp)

	// create the hls manifest

	hlsManifest := &models.HLSManifest{
		Name:         stringToPtr("my-awesome-hls-manifest"),
		ManifestName: stringToPtr("manifest.m3u8"),
		Outputs:      []models.Output{manifestOutput},
	}

	hlsService := services.NewHLSManifestService(bitmovin)

	hlsManifestResp, err := hlsService.Create(hlsManifest)
	errorHandler(hlsManifestResp.Status, err)

	// add the audio media and variant stream

	audioMediaInfo := &models.MediaInfo{
		Type:        bitmovintypes.MediaTypeAudio,
		URI:         stringToPtr("audiomedia.m3u8"),
		GroupID:     stringToPtr("audio_group"),
		Language:    stringToPtr("en"),
		Name:        stringToPtr("HLS Audio Media"),
		SegmentPath: stringToPtr(outputPath + "/audio/ts"),
		EncodingID:  encodingResp.Data.Result.ID,
		StreamID:    aacStreamResp.Data.Result.ID,
		MuxingID:    audioTSMuxingResp.Data.Result.ID,
	}
	audioMediaInfoResp, err := hlsService.AddMediaInfo(*hlsManifestResp.Data.Result.ID, audioMediaInfo)
	errorHandler(audioMediaInfoResp.Status, err)

	video1080pStreamInfo := &models.StreamInfo{
		Audio:       stringToPtr("audio_group"),
		SegmentPath: stringToPtr(outputPath + "/video/1080p/ts"),
		URI:         stringToPtr("video_1080p.m3u8"),
		EncodingID:  encodingResp.Data.Result.ID,
		StreamID:    videoStream1080pResp.Data.Result.ID,
		MuxingID:    videoTSMuxing1080pResp.Data.Result.ID,
	}
	video1080pStreamInfoResponse, err := hlsService.AddStreamInfo(*hlsManifestResp.Data.Result.ID, video1080pStreamInfo)
	errorHandler(video1080pStreamInfoResponse.Status, err)

	startResp, err = hlsService.Start(*hlsManifestResp.Data.Result.ID)
	errorHandler(startResp.Status, err)

	waitForHlsManifest(hlsService, hlsManifestResp)
}

func waitForEncoding(encodingS *services.EncodingService, encodingResp *models.EncodingResponse) {
	status := ""

	for status != "FINISHED" {
		statusResp, err := encodingS.RetrieveStatus(*encodingResp.Data.Result.ID)
		if err != nil {
			fmt.Println("Error in encoding status")
			return
		}
		status = *statusResp.Data.Result.Status
		fmt.Printf("Encoding Status: %s\n", status)

		if status == "ERROR" {
			fmt.Println("Error in encoding status")
			return
		}

		time.Sleep(10 * time.Second)
	}
}

func waitForDashManifest(dashService *services.DashManifestService, dashManifestResp *models.DashManifestResponse) {
	status := ""

	for status != "FINISHED" {
		statusResp, err := dashService.RetrieveStatus(*dashManifestResp.Data.Result.ID)
		if err != nil {
			fmt.Println("Error in manifest status")
			return
		}
		status = *statusResp.Data.Result.Status
		fmt.Printf("DASH manifest Status: %s\n", status)

		if status == "ERROR" {
			fmt.Println("Error in manifest status")
			return
		}

		time.Sleep(10 * time.Second)
	}
}

func waitForHlsManifest(hlsService *services.HLSManifestService, hlsManifestResp *models.HLSManifestResponse) {
	status := ""

	for status != "FINISHED" {
		statusResp, err := hlsService.RetrieveStatus(*hlsManifestResp.Data.Result.ID)
		if err != nil {
			fmt.Println("Error in manifest status")
			return
		}
		status = *statusResp.Data.Result.Status
		fmt.Printf("HLS manifest Status: %s\n", status)

		if status == "ERROR" {
			fmt.Println("Error in manifest status")
			return
		}

		time.Sleep(10 * time.Second)
	}
}

func errorHandler(responseStatus bitmovintypes.ResponseStatus, err error) {
	if err != nil {
		fmt.Println("go error")
		fmt.Println(err)
	} else if responseStatus == "ERROR" {
		fmt.Println("api error")
	}
}

func stringToPtr(s string) *string {
	return &s
}

func intToPtr(i int64) *int64 {
	return &i
}

func boolToPtr(b bool) *bool {
	return &b
}

func floatToPtr(f float64) *float64 {
	return &f
}
