# vim: set expandtab sw=4
FROM golang:1.14-buster

# Create a non-priviledged user to run browsers as (Firefox and Chrome do not
# like to run as root).
RUN chmod a+rx $HOME && useradd --uid 9999 --user-group --create-home browser

# Sort the package names!
# firefox-esr: provides deps for Firefox (we don't use ESR directly)
# openjdk-11-jdk: provides JDK/JRE to Selenium & gcloud SDK
# python-crcmod: native module to speed up CRC checksum in gsutil
RUN apt-get update -qqy && apt-get install -qqy --no-install-suggests \
        curl \
        firefox-esr \
        lsb-release \
        openjdk-11-jdk \
        python-crcmod \
        python3.7 \
        sudo \
        tox \
        wget \
        xvfb && \
    rm /usr/bin/firefox

# The base golang image adds Go paths to PATH, which cannot be inherited in
# sudo by default because of the `secure_path` directive. Overwrite sudoers to
# discard the setting.
RUN echo "root ALL=(ALL:ALL) ALL" > /etc/sudoers

# Node LTS
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash - && \
    apt-get install -qqy nodejs

# Google Cloud SDK
# Based on https://github.com/GoogleCloudPlatform/cloud-sdk-docker/blob/master/Dockerfile
RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb https://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" > /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update -qqy && apt-get install -qqy \
        google-cloud-sdk \
        google-cloud-sdk-app-engine-python \
        google-cloud-sdk-app-engine-python-extras \
        google-cloud-sdk-app-engine-go \
        google-cloud-sdk-datastore-emulator && \
    gcloud config set core/disable_usage_reporting true && \
    gcloud config set component_manager/disable_update_check true && \
    gcloud --version
