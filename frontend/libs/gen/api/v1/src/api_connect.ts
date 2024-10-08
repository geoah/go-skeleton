// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts,import_extension=none"
// @generated from file api.proto (package api.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PingRequest, PingResponse } from "./api_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.v1.APIService
 */
export const APIService = {
  typeName: "api.v1.APIService",
  methods: {
    /**
     * @generated from rpc api.v1.APIService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

