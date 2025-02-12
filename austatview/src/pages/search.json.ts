import {getCollection} from "astro:content";

async function getPosts(){
  const dset = (await getCollection('metadata'))
  return dset;
}

export async function GET({}) {
    return new Response (JSON.stringify(await getPosts
    ()), {
      status: 200,
      headers: {
        "Content-Type": "application/json",
      },
    });
}
