# use the official Bun image
FROM oven/bun:latest AS base
WORKDIR /app

# install dependencies
COPY package.json bun.lockb ./
# test the lockfile, if do not find exit
RUN bun install --frozen-lockfile

# copy the rest of the application code
COPY . .

ENV NODE_ENV production
RUN bun run build

# run the application
USER bun
EXPOSE 3000
CMD ["bun","run","start"]



