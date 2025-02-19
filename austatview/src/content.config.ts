import { glob,file } from 'astro/loaders';
import { defineCollection, z } from 'astro:content';

const featured = defineCollection({
	// Load Markdown and MDX files in the `src/content/blog/` directory.
	loader: glob({ base: './src/content/featured', pattern: '**/*.{md,mdx}' }),
	// Type-check frontmatter using a schema
	schema: z.object({
		title: z.string(),
		description: z.string(),
		// Transform string to Date object
		pubDate: z.coerce.date(),
		updatedDate: z.coerce.date().optional(),
		heroImage: z.string().optional(),
	}),
});

const metadata = defineCollection({
  loader: file('./src/content/metadata/reference_stubs.json',
               {parser: (text) => JSON.parse(text).data.dataflows}),
})

export const collections = { featured ,
  metadata
};
