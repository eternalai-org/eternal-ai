package services

import (
	"context"
	"fmt"
)

func (s *Service) GetVideoMagicPromptFromImage(ctx context.Context, userPrompt string, imageUrl string) (string, error) {
	descriptionImage, err := s.GetDescriptionImage(ctx, imageUrl)
	if err != nil {
		return "", err
	}
	input := map[string]interface{}{
		"model": "Llama3.3",
		"messages": []map[string]interface{}{
			map[string]interface{}{
				"role":    "system",
				"content": "You are a helpful assistant specializing in creating smooth and engaging short video content.",
			},
			map[string]interface{}{
				"role":    "user",
				"content": "Describe this image.",
			},
			map[string]interface{}{
				"role":    "assistant",
				"content": descriptionImage,
			},
			map[string]interface{}{
				"role":    "user",
				"content": fmt.Sprintf("I would like to create a short video based on my request: \\n%v \\n Prioritize my request, even if it contrasts with the content of the image. For example, if the image features characters laughing and having fun, my request should be to create a video where they are crying or fighting. The video should exclusively show them crying or fighting. Just provide the prompt.", userPrompt),
			},
		},
	}
	return s.openais["Agent"].CallDirectlyEternalLLMV2(input, s.conf.KnowledgeBaseConfig.DirectServiceUrl)
}

func (s *Service) GetDescriptionImage(ctx context.Context, imageUrl string) (string, error) {
	input := map[string]interface{}{
		"model": "Qwen/Qwen2.5-VL-7B-Instruct",
		"messages": []map[string]interface{}{
			map[string]interface{}{
				"role": "system",
				"content": []map[string]interface{}{
					map[string]interface{}{
						"type": "text",
						"text": "You are a prompt optimization specialist whose goal is to rewrite the user's input prompts into high-quality English prompts by referring to the details of the user's input images, making them more complete and expressive while maintaining the original meaning. You need to integrate the content of the user's photo with the input prompt for the rewrite, strictly adhering to the formatting of the examples provided.\\nTask Requirements:\\n1. For overly brief user inputs, reasonably infer and supplement details without changing the original meaning, making the image more complete and visually appealing;\\n2. Improve the characteristics of the main subject in the user's description (such as appearance, expression, quantity, ethnicity, posture, etc.), rendering style, spatial relationships, and camera angles;\\n3. The overall output should be in Chinese, retaining original text in quotes and book titles as well as important input information without rewriting them;\\n4. The prompt should match the user’s intent and provide a precise and detailed style description. If the user has not specified a style, you need to carefully analyze the style of the user's provided photo and use that as a reference for rewriting;\\n5. If the prompt is an ancient poem, classical Chinese elements should be emphasized in the generated prompt, avoiding references to Western, modern, or foreign scenes;\\n6. You need to emphasize movement information in the input and different camera angles;\\n7. Your output should convey natural movement attributes, incorporating natural actions related to the described subject category, using simple and direct verbs as much as possible;\\n8. You should reference the detailed information in the image, such as character actions, clothing, backgrounds, and emphasize the details in the photo;\\n9. Control the rewritten prompt to around 80-100 words.\\n10. No matter what language the user inputs, you must always output in English.\\nExample of the rewritten English prompt:\\n1. A Japanese fresh film-style photo of a young East Asian girl with double braids sitting by the boat. The girl wears a white square collar puff sleeve dress, decorated with pleats and buttons. She has fair skin, delicate features, and slightly melancholic eyes, staring directly at the camera. Her hair falls naturally, with bangs covering part of her forehead. She rests her hands on the boat, appearing natural and relaxed. The background features a blurred outdoor scene, with hints of blue sky, mountains, and some dry plants. The photo has a vintage film texture. A medium shot of a seated portrait.\\n2. An anime illustration in vibrant thick painting style of a white girl with cat ears holding a folder, showing a slightly dissatisfied expression. She has long dark purple hair and red eyes, wearing a dark gray skirt and a light gray top with a white waist tie and a name tag in bold Chinese characters that says \\\"紫阳\\\" (Ziyang). The background has a light yellow indoor tone, with faint outlines of some furniture visible. A pink halo hovers above her head, in a smooth Japanese cel-shading style. A close-up shot from a slightly elevated perspective.\\n3. CG game concept digital art featuring a huge crocodile with its mouth wide open, with trees and thorns growing on its back. The crocodile's skin is rough and grayish-white, resembling stone or wood texture. Its back is lush with trees, shrubs, and thorny protrusions. With its mouth agape, the crocodile reveals a pink tongue and sharp teeth. The background features a dusk sky with some distant trees, giving the overall scene a dark and cold atmosphere. A close-up from a low angle.\\n4. In the style of an American drama promotional poster, Walter White sits in a metal folding chair wearing a yellow protective suit, with the words \\\"Breaking Bad\\\" written in sans-serif English above him, surrounded by piles of dollar bills and blue plastic storage boxes. He wears glasses, staring forward, dressed in a yellow jumpsuit, with his hands resting on his knees, exuding a calm and confident demeanor. The background shows an abandoned, dim factory with light filtering through the windows. There’s a noticeable grainy texture. A medium shot with a straight-on close-up of the character.\\nDirectly output the rewritten English text.",
					},
				},
			},
			map[string]interface{}{
				"role": "user",
				"content": []map[string]interface{}{
					map[string]interface{}{
						"type": "text",
						"text": "Describe this image.",
					},
					map[string]interface{}{
						"type": "image_url",
						"image_url": map[string]interface{}{
							"url": imageUrl,
						},
					},
				},
			},
		},
	}
	return s.openais["Agent"].CallDirectlyEternalLLMV2(input, s.conf.KnowledgeBaseConfig.DirectServiceUrl)
}
