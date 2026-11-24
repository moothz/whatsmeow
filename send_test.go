// Copyright (c) 2026 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package whatsmeow

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow/proto/waE2E"
)

func TestGetTypeAndMediaTypeAudioStickerMessage(t *testing.T) {
	msg := &waE2E.Message{
		AudioStickerMessage: &waE2E.FutureProofMessage{
			Message: &waE2E.Message{
				StickerMessage: &waE2E.StickerMessage{
					Mimetype: proto.String("image/webp"),
					Audio: &waE2E.StickerMessage_AudioMessage{
						AudioMessage: &waE2E.AudioMessage{
							Mimetype: proto.String("audio/ogg; codecs=opus"),
						},
					},
				},
			},
		},
	}

	if got := getTypeFromMessage(msg); got != "media" {
		t.Fatalf("getTypeFromMessage(AudioStickerMessage) = %q, want %q", got, "media")
	}
	if got := getMediaTypeFromMessage(msg); got != "sticker" {
		t.Fatalf("getMediaTypeFromMessage(AudioStickerMessage) = %q, want %q", got, "sticker")
	}
}
