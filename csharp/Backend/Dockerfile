FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build

WORKDIR /src

# Copy the Backend and Database project files and restore dependencies
COPY Backend/Backend.csproj Backend/

# Restore dependencies
RUN dotnet restore ./Backend/Backend.csproj

# Copy the entire source code
COPY . .

WORKDIR /src/Backend
RUN dotnet publish Backend.csproj -c Release -o /app

FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS base

WORKDIR /app

COPY --from=build /app .

ENV ASPNETCORE_URLS=http://+:${BACKEND_HTTP_PORT}
ENV ASPNETCORE_ENVIRONMENT=Development


ENTRYPOINT ["dotnet", "Backend.dll"]
