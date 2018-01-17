/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package eventemitter

const (
	OrderCanceled                  = "OrderCanceled"
	OrderFilled                    = "OrderFilled"
	ExtractorFork                  = "ExtractorFork" //chain forked
	OrderManagerFork               = "OrderManagerFork"
	RingSubmitFailed               = "RingSubmitFailed" //submit ring failed
	Transaction                    = "Transaction"
	Gateway                        = "Gateway"
	AccountTransfer                = "AccountTransfer"
	AccountApproval                = "AccountApproval"
	TokenRegistered                = "TokenRegistered"
	TokenUnRegistered              = "TokenUnRegistered"
	RingHashSubmitted              = "RingHashSubmitted"
	AddressAuthorized              = "AddressAuthorized"
	AddressDeAuthorized            = "AddressDeAuthorized"
	OrderManagerGatewayNewOrder    = "OrderManagerGatewayNewOrder"
	OrderManagerExtractorRingMined = "OrderManagerRingMined"
	OrderManagerExtractorFill      = "OrderManagerExtractorFill"
	OrderManagerExtractorCancel    = "OrderManagerExtractorCancel"
	OrderManagerExtractorCutoff    = "OrderManagerExtractorCutoff"
	MinedOrderState                = "MinedOrderState" //orderbook send orderstate to miner

	//Miner
	MinerDeleteOrderState          = "MinerDeleteOrderState"
	MinerNewOrderState             = "MinerNewOrderState"
	MinerNewRing                   = "MinerNewRing"
	MinerRingMined                 = "MinerRingMined"
	MinerRingSubmitFailed          = "MinerRingSubmitFailed"
	MinerSubmitRingMethod          = "MinerSubmitRingMethod"
	MinerSubmitRingHashMethod      = "MinerSubmitRingHashMethod"
	MinerBatchSubmitRingHashMethod = "MinerBatchSubmitRingHashMethod"

	// Block
	BlockNew = "BlockNew"

	// Extractor
	SyncChainComplete = "SyncChainComplete"
	ChainForkDetected = "ChainForkDetected"
	ChainForkProcess  = "ChainForkProcess"

	// todo: Methods
	//WethDepositMethod    = "WethDepositMethod"
	//WethWithdrawalMethod = "WethWithdrawalMethod"
	//ApproveMethod        = "ApproveMethod"
)

//var watchers map[string][]*Watcher
//var mtx *sync.Mutex

type EventData interface{}

type Watcher struct {
	Concurrent bool
	Handle     func(eventData EventData) error
}

type Emitter interface {
	Un(topic string, watcher *Watcher)
	On(topic string, watcher *Watcher)
	Emit(topic string, eventData EventData)
}
