/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package server implements a server for Greeter service.
package server

import (
	"context"
	"librarium/api"
	"librarium/database_access"
	"log"
)

//Server ...
type Server struct{}

// GetAuthor ...
func (s *Server) GetAuthor(ctx context.Context, req *api.DataRequest) (*api.DataReply, error) {
	log.Printf(`Запрошен автор книги "%s"`, req.AskMessage)
	return &api.DataReply{ReplyMessage: database_access.ReadBookAuthor(req.AskMessage)}, nil
}

func (s *Server) GetBooks(ctx context.Context, req *api.DataRequest) (*api.DataReply, error) {
	log.Printf(`Запрошены книги автора "%s"`, req.AskMessage)
	return &api.DataReply{ReplyMessage: database_access.ReadAuthorBook(req.AskMessage)}, nil
}
