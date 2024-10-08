# SPDX-FileCopyrightText: 2023 Siemens AG
#
# SPDX-License-Identifier: Apache-2.0
#
# Author: Michael Adler <michael.adler@siemens.com>
FROM gcr.io/distroless/static-debian12:nonroot@sha256:bc539c0f2a6047152f4b919c04be587fcb935a1e550dac7debe94cff0be02ddb

# this file is generated by goreleaser as part of the CI pipeline
COPY wfx /usr/bin/wfx

EXPOSE 8080 8081

ENTRYPOINT ["wfx"]
