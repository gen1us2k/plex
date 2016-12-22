package plex

import "encoding/xml"

type (
	PMSServer struct {
		AccessToken       string `xml:"accessToken,attr"`
		Name              string `xml:"name,attr"`
		Address           string `xml:"address,attr"`
		Port              string `xml:"port,attr"`
		Version           string `xml:"version,attr"`
		Scheme            string `xml:"scheme,attr"`
		Host              string `xml:"host,attr"`
		LocalAddresses    string `xml:"localAddresses,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		CreatedAt         string `xml:"createdAt,attr"`
		UpdatedAt         string `xml:"updatedAt,attr"`
		Owned             string `xml:"owned,attr"`
		Synced            string `xml:"synced,attr"`
	}
	PMSPlaylist struct {
		RatingKey    string `xml:"ratingKey,attr"`
		Key          string `xml:"key,attr"`
		Guid         string `xml:"guid,attr"`
		Type         string `xml:"type,attr"`
		Title        string `xml:"title,attr"`
		Summary      string `xml:"summary,attr"`
		Smart        string `xml:"smart,attr"`
		PlaylistType string `xml:"playlistType,attr"`
		Composite    string `xml:"composite,attr"`
		ViewCount    string `xml:"viewCount,attr"`
		LastViewed   string `xml:"lastViewedAt,attr"`
		Duration     string `xml:"duration,attr"`
		LeafCount    string `xml:"leafCount,attr"`
		Added        string `xml:"addedAt,attr"`
		Updated      string `xml:"updatedAt,attr"`
	}
	PMSPlaylistInfo struct {
		XMLName xml.Name `xml:"MediaContainer"`
		Videos  []Video  `xml:"Video"`
	}
	ServersResponse struct {
		XMLName xml.Name    `xml:"MediaContainer"`
		Servers []PMSServer `xml:"Server"`
	}
	Video struct {
		RatingKey            string `json:"ratingKey" xml:"ratingKey,attr"`
		Key                  string `json:"key" xml:"key,attr"`
		GrandparentTitle     string `json:"grandparentTitle" xml:"grandparentTitle,attr"`
		GrandparentRatingKey string `json:"grandparentRatingKey" xml:"grandparentRatingKey,attr"`
		ParentRatingKey      string `json:"parentRatingKey" xml:"parentRatingKey,attr"`
		ParentYear           string `json:"parentYear" xml:"parentYear,attr"`
		ParentTitle          string `json:"parentTitle" xml:"parentTitle,attr"`
		GUID                 string `json:"guid" xml:"guid,attr"`
		LibrarySectionID     string `json:"librarySectionID" xml:"librarySectionID,attr"`
		Type                 string `json:"type" xml:"type,attr"`
		Title                string `json:"title" xml:"title,attr"`
		Summary              string `json:"summary" xml:"summary,attr"`
		ViewCount            string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt         string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year                 string `json:"year" xml:"year,attr"`
		Thumb                string `json:"thumb" xml:"thumb,attr"`
		Art                  string `json:"art" xml:"art,attr"`
		Duration             string `json:"duration" xml:"duration,attr"`
		AddedAt              string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt            string `json:"updatedAt" xml:"updatedAt,attr"`
		ChapterSource        string `json:"chapterSource" xml:"chapterSource,attr"`
		Media                struct {
			VideoResolution string `json:"videoResolution" xml:"videoResolution,attr"`
			ID              string `json:"id" xml:"id,attr"`
			Duration        string `json:"duration" xml:"duration,attr"`
			Bitrate         string `json:"bitrate" xml:"bitrate,attr"`
			Width           string `json:"width" xml:"width,attr"`
			Height          string `json:"height" xml:"height,attr"`
			AspectRatio     string `json:"aspectRatio" xml:"aspectRatio,attr"`
			AudioChannels   string `json:"audioChannels" xml:"audioChannels,attr"`
			AudioCodec      string `json:"audioCodec" xml:"audioCodec,attr"`
			VideoCodec      string `json:"videoCodec" xml:"videoCodec,attr"`
			Container       string `json:"container" xml:"container,attr"`
			VideoFrameRate  string `json:"videoFrameRate" xml:"videoFrameRate,attr"`
			VideoProfile    string `json:"videoProfile" xml:"videoProfile,attr"`
			Part            struct {
				ID           string `json:"id" xml:"id,attr"`
				Key          string `json:"key" xml:"key,attr"`
				Duration     string `json:"duration" xml:"duration,attr"`
				File         string `json:"file" xml:"file,attr"`
				Size         string `json:"size" xml:"size,attr"`
				Container    string `json:"container" xml:"container,attr"`
				VideoProfile string `json:"videoProfile" xml:"videoProfile,attr"`
				Stream       []struct {
					ID                 string `json:"id" xml:"id,attr"`
					StreamType         string `json:"streamType" xml:"streamType,attr"`
					Default            string `json:"default" xml:"default,attr"`
					Codec              string `json:"codec" xml:"codec,attr"`
					Index              string `json:"index" xml:"index,attr"`
					Bitrate            string `json:"bitrate" xml:"bitrate,attr"`
					BitDepth           string `json:"bitDepth" xml:"bitDepth,attr"`
					Cabac              string `json:"cabac" xml:"cabac,attr"`
					ChromaSubsampling  string `json:"chromaSubsampling" xml:"chromaSubsampling,attr"`
					CodecID            string `json:"codecID" xml:"codecID,attr"`
					ColorRange         string `json:"colorRange" xml:"colorRange,attr"`
					ColorSpace         string `json:"colorSpace" xml:"colorSpace,attr"`
					Duration           string `json:"duration" xml:"duration,attr"`
					FrameRate          string `json:"frameRate" xml:"frameRate,attr"`
					FrameRateMode      string `json:"frameRateMode" xml:"frameRateMode,attr"`
					HasScalingMatrix   string `json:"hasScalingMatrix" xml:"hasScalingMatrix,attr"`
					HeaderStripping    string `json:"headerStripping" xml:"headerStripping,attr"`
					Height             string `json:"height" xml:"height,attr"`
					Level              string `json:"level" xml:"level,attr"`
					PixelFormat        string `json:"pixelFormat" xml:"pixelFormat,attr"`
					Profile            string `json:"profile" xml:"profile,attr"`
					RefFrames          string `json:"refFrames" xml:"refFrames,attr"`
					ScanType           string `json:"scanType" xml:"scanType,attr"`
					Width              string `json:"width" xml:"width,attr"`
					Selected           string `json:"selected" xml:"selected,attr"`
					Channels           string `json:"channels" xml:"channels,attr"`
					AudioChannelLayout string `json:"audioChannelLayout" xml:"audioChannelLayout,attr"`
					BitrateMode        string `json:"bitrateMode" xml:"bitrateMode,attr"`
					DialogNorm         string `json:"dialogNorm" xml:"dialogNorm,attr"`
					SamplingRate       string `json:"samplingRate" xml:"samplingRate,attr"`
				} `json:"stream" xml:"Stream"`
			} `json:"part" xml:"Part"`
		} `json:"media" xml:"Media"`
	}
	PlaylistResponse struct {
		XMLName   xml.Name      `xml:"MediaContainer"`
		Playlists []PMSPlaylist `xml:"Playlist"`
	}
	headers struct {
		Platform         string
		PlatformVersion  string
		Provides         string
		ClientIdentifier string
		Product          string
		Version          string
		Device           string
		ContainerSize    string
		ContainerStart   string
		Token            string
		Accept           string
		ContentType      string
	}
)
